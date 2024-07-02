import axios from 'axios';

axios.defaults.headers.common['Content-Type'] = 'application/json'
axios.defaults.withCredentials = true

axios.interceptors.request.use(config => {
    const token = localStorage.getItem('jwtToken')
    if (token) {
        config.headers.Authorization = `Bearer ${token}`
    }
    return config;
}, error => Promise.reject(error))

axios.interceptors.response.use(response => {
    if (response.data && response.data.token) {
        localStorage.setItem('jwtToken', response.data.token)
    }
    return response
}, error => {
    if (error.response && error.response.status === 401) {
        // Token expired or invalid
        console.error('Token expired or invalid.');
        // Optional: Redirect to login
    }
    return Promise.reject(error)
})

export default axios
