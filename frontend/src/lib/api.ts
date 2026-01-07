import axios from 'axios'

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

export const api = axios.create({
  baseURL: `${API_URL}/api`,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Auth endpoints
export const authAPI = {
  getStatus: () => api.get('/auth/status'),
  spotifyLogin: () => window.location.href = `${API_URL}/api/auth/spotify`,
  appleLogin: () => window.location.href = `${API_URL}/api/auth/apple`,
}

// Migration endpoints
export const migrationAPI = {
  start: (data: {
    source_provider: string
    target_provider: string
    playlist_ids?: string[]
    migrate_library?: boolean
  }) => api.post('/migrations', data),
  
  getStatus: (id: string) => api.get(`/migrations/${id}`),
  
  list: () => api.get('/migrations'),
}

// Playlist endpoints
export const playlistAPI = {
  getAll: (provider: string) => api.get(`/playlists/${provider}`),
  
  getTracks: (provider: string, playlistId: string) => 
    api.get(`/playlists/${provider}/${playlistId}`),
}

export default api
