import request from '@/utils/request'

// 用户登录
export function login (data) {
  return request({
    url: '/login',
    method: 'post',
    data
  })
}

// 用户注册
export function register (data) {
  return request({
    url: '/register',
    method: 'post',
    data
  })
}

// 获取下级用户列表
export function getSubUsers (userRole) {
  return request({
    url: '/users/sub',
    method: 'get',
    params: { role: userRole }
  })
}

// 创建下级用户
export function createSubUser (userRole, data) {
  return request({
    url: '/users/create',
    method: 'post',
    data: {
      ...data,
      role: userRole
    }
  })
}
