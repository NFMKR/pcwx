import { createStore } from 'vuex'
import { login, logout } from '../api/auth'

const store = createStore({
  state: {
    token: localStorage.getItem('token') || '',
    user: JSON.parse(localStorage.getItem('user')) || null
  },

  mutations: {
    SET_TOKEN (state, token) {
      state.token = token
      localStorage.setItem('token', token)
    },
    SET_USER (state, user) {
      state.user = user
      localStorage.setItem('user', JSON.stringify(user))
    },
    CLEAR_USER_INFO (state) {
      state.token = ''
      state.user = null
      localStorage.removeItem('token')
      localStorage.removeItem('user')
    }
  },

  actions: {
    // 用户登录
    async login ({ commit }, userInfo) {
      try {
        const response = await login(userInfo)
        console.log('Login API Response:', response)

        if (!response || !response.token) {
          throw new Error('登录失败：未获取到token')
        }

        // 保存token和用户信息
        commit('SET_TOKEN', response.token)
        commit('SET_USER', response.data || {
          username: userInfo.Username,
          role: 'admin'
        })

        return response
      } catch (error) {
        console.error('Login Error:', error)
        throw error
      }
    },

    // 用户登出
    async logout ({ commit }) {
      try {
        await logout()
        commit('CLEAR_USER_INFO')
      } catch (error) {
        console.error('Logout Error:', error)
        throw error
      }
    }
  },

  getters: {
    isAuthenticated: state => !!state.token,
    userRole: state => state.user?.role
  }
})

export default store
