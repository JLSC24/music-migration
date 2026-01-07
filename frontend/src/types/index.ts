export interface Track {
  id: string
  name: string
  artists: string[]
  album: string
  isrc?: string
  duration_ms: number
}

export interface Playlist {
  id: string
  name: string
  description?: string
  track_count: number
  tracks?: Track[]
}

export interface Migration {
  id: number
  user_id: number
  source_provider: string
  target_provider: string
  status: 'pending' | 'processing' | 'completed' | 'failed'
  total_tracks: number
  matched_tracks: number
  failed_tracks: number
  created_at: string
  completed_at?: string
}

export interface AuthStatus {
  spotify: boolean
  apple_music: boolean
}
