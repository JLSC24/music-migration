'use client'

import { useState } from 'react'
import Link from 'next/link'
import { ArrowLeft, Music, Loader2 } from 'lucide-react'

export default function MigratePage() {
  const [source, setSource] = useState<'spotify' | 'apple_music'>('spotify')
  const [target, setTarget] = useState<'spotify' | 'apple_music'>('apple_music')
  const [migrating, setMigrating] = useState(false)
  const [progress, setProgress] = useState(0)

  const handleMigrate = async () => {
    setMigrating(true)
    // TODO: Implement migration logic
    
    // Simulate progress
    const interval = setInterval(() => {
      setProgress(prev => {
        if (prev >= 100) {
          clearInterval(interval)
          setMigrating(false)
          return 100
        }
        return prev + 10
      })
    }, 500)
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="container mx-auto px-4 py-8">
        <Link href="/dashboard" className="inline-flex items-center text-purple-600 hover:text-purple-700 mb-6">
          <ArrowLeft className="w-4 h-4 mr-2" />
          Back to Dashboard
        </Link>

        <div className="max-w-2xl mx-auto">
          <div className="bg-white rounded-xl shadow-sm border p-8">
            <h1 className="text-3xl font-bold mb-8">Migrate Your Music</h1>

            {/* Source/Target Selection */}
            <div className="space-y-6 mb-8">
              <div>
                <label className="block text-sm font-medium mb-2">From</label>
                <select
                  value={source}
                  onChange={(e) => setSource(e.target.value as any)}
                  className="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent"
                >
                  <option value="spotify">Spotify</option>
                  <option value="apple_music">Apple Music</option>
                </select>
              </div>

              <div className="text-center">
                <div className="inline-block p-2 bg-purple-100 rounded-full">
                  <svg className="w-6 h-6 text-purple-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 14l-7 7m0 0l-7-7m7 7V3" />
                  </svg>
                </div>
              </div>

              <div>
                <label className="block text-sm font-medium mb-2">To</label>
                <select
                  value={target}
                  onChange={(e) => setTarget(e.target.value as any)}
                  className="w-full px-4 py-2 border rounded-lg focus:ring-2 focus:ring-purple-500 focus:border-transparent"
                >
                  <option value="spotify">Spotify</option>
                  <option value="apple_music">Apple Music</option>
                </select>
              </div>
            </div>

            {/* Migration Options */}
            <div className="mb-8">
              <h3 className="font-semibold mb-3">What to migrate?</h3>
              <div className="space-y-2">
                <label className="flex items-center">
                  <input type="checkbox" className="mr-2 rounded" defaultChecked />
                  <span>All Playlists</span>
                </label>
                <label className="flex items-center">
                  <input type="checkbox" className="mr-2 rounded" defaultChecked />
                  <span>Liked Songs</span>
                </label>
                <label className="flex items-center">
                  <input type="checkbox" className="mr-2 rounded" />
                  <span>Albums</span>
                </label>
              </div>
            </div>

            {/* Progress Bar */}
            {migrating && (
              <div className="mb-6">
                <div className="flex items-center justify-between mb-2">
                  <span className="text-sm font-medium">Migrating...</span>
                  <span className="text-sm font-medium">{progress}%</span>
                </div>
                <div className="w-full bg-gray-200 rounded-full h-2">
                  <div
                    className="bg-purple-600 h-2 rounded-full transition-all duration-300"
                    style={{ width: `${progress}%` }}
                  />
                </div>
              </div>
            )}

            {/* Action Button */}
            <button
              onClick={handleMigrate}
              disabled={migrating || source === target}
              className="w-full px-6 py-3 bg-purple-600 hover:bg-purple-700 disabled:bg-gray-400 text-white font-semibold rounded-lg transition-colors flex items-center justify-center"
            >
              {migrating ? (
                <>
                  <Loader2 className="w-5 h-5 mr-2 animate-spin" />
                  Migrating...
                </>
              ) : (
                <>
                  <Music className="w-5 h-5 mr-2" />
                  Start Migration
                </>
              )}
            </button>

            {source === target && (
              <p className="text-red-500 text-sm mt-2 text-center">
                Source and target must be different
              </p>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}
