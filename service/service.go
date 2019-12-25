package service

import (
	"fmt"
	"sync"

	"miMallDemo/model"
	"miMallDemo/utils"

)

// 一般在 handler 中主要做解析参数、返回数据操作，简单的逻辑也可以在 handler 中做，
// 像新增用户、删除用户、更新用户，代码量不大，所以也可以放在 handler 中。
// 有些代码量很大的逻辑就不适合放在 handler 中，
// 因为这样会导致 handler 逻辑不是很清晰，
// 这时候实际处理的部分通常放在 service 包中。

func ListUser(username string, offset, limit int) ([]*model.UserInfo, uint, error) {

	infos := make([]*model.UserInfo, 0)	// json用户信息格式的数组
	users, count, err := model.ListUser(username, offset, limit) // 数据库中获取用户信息；按照 username 查询
	if err != nil {
		return nil, count, err
	}

	var ids []uint	// 用户ID数组
	for _, user := range users {
		ids = append(ids, user.ID)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock: new(sync.Mutex),
		IdMap: make(map[uint]*model.UserInfo, len(users)),
	} // 带锁的用户列表

	errChan := make(chan error, 1) // 通道，传递错误信息，零值是nil
	finished := make(chan bool, 1) // 通道，传递是否结束，零值是false

	for _, u := range users {
		wg.Add(1)
		go func(u *model.User) {
			defer wg.Done()

			shortId, err := utils.GetShortId()
			if err != nil {
				errChan <- err
			}

			userList.Lock.Lock()	// 加锁
			defer userList.Lock.Unlock()	// 解锁
			userList.IdMap[u.ID] = &model.UserInfo{
				ID:			u.ID,
				Username:	u.Username,
				SayHello:  	fmt.Sprintf("Hello %s", shortId),
				Password:  	u.Password,
				CreatedAt: 	u.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: 	u.CreatedAt.Format("2006-01-02 15:04:05"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}

	return infos, count, nil

}