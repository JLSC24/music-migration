'use client'

import { useEffect, useState } from 'react'
import Link from 'next/link'
import { Music, LogOut, Settings } from 'lucide-react'

export default function DashboardPage() {
  const [authStatus, setAuthStatus] = useState({
    spotify: false,
    apple_music: false
  })

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm border-b">
        <div className="container mx-auto px-4 py-4">
          <div className="flex items-center justify-between">
            <Link href="/" className="flex items-center gap-2">
              <Music className="w-8 h-8 text-purple-600" />
              <span className="text-xl font-bold">Music Migration</span>
            </Link>
            <div className="flex items-center gap-4">
              <button className="p-2 hover:bg-gray-100 rounded-lg">
                <Settings className="w-5 h-5" />
              </button>
              <button className="p-2 hover:bg-gray-100 rounded-lg">
                <LogOut className="w-5 h-5" />
              </button>
            </div>
          </div>
        </div>
      </header>

      <div className="container mx-auto px-4 py-8">
        <h1 className="text-3xl font-bold mb-8">Dashboard</h1>

        {/* Auth Status Cards */}
        <div className="grid md:grid-cols-2 gap-6 mb-8">
          <AuthCard
            provider="Spotify"
            connected={authStatus.spotify}
            color="bg-green-500"
            onConnect={() => window.location.href = `${process.env.NEXT_PUBLIC_API_URL}/api/auth/spotify`}
          />
          <AuthCard
            provider="Apple Music"
            connected={authStatus.apple_music}
            color="bg-red-500"
            onConnect={() => window.location.href = `${process.env.NEXT_PUBLIC_API_URL}/api/auth/apple`}
          />
        </div>

        {/* Migration Section */}
        {authStatus.spotify && authStatus.apple_music ? (
          <div className="bg-white rounded-xl shadow-sm border p-8">
            <h2 className="text-2xl font-bold mb-6">Start Migration</h2>
            <Link
              href="/migrate"
              className="inline-block px-6 py-3 bg-purple-600 hover:bg-purple-700 text-white font-semibold rounded-lg transition-colors"
            >
              Start New Migration
            </Link>
          </div>
        ) : (
          <div className="bg-yellow-50 border border-yellow-200 rounded-xl p-6">
            <p className="text-yellow-800">
              Please connect both Spotify and Apple Music accounts to start migrating your music.
            </p>
          </div>
        )}

        {/* Recent Migrations */}
        <div className="mt-8">
          <h2 className="text-2xl font-bold mb-4">Recent Migrations</h2>
          <div className="bg-white rounded-xl shadow-sm border p-6">
            <p className="text-gray-500 text-center py-8">No migrations yet</p>
          </div>
        </div>
      </div>
    </div>
  )
}

function AuthCard({ 
  provider, 
  connected, 
  color,
  onConnect 
}: { 
  provider: string
  connected: boolean
  color: string
  onConnect: () => void
}) {
  return (
    <div className="bg-white rounded-xl shadow-sm border p-6">
      <div className="flex items-center justify-between mb-4">
        <h3 className="text-xl font-semibold">{provider}</h3>
        <div className={`w-3 h-3 rounded-full ${connected ? 'bg-green-500' : 'bg-gray-300'}`} />
      </div>
      <p className="text-gray-600 mb-4">
        {connected ? 'Connected' : 'Not connected'}
      </p>
      {!connected && (
        <button
          onClick={onConnect}
          className={`px-4 py-2 ${color} text-white font-semibold rounded-lg hover:opacity-90 transition-opacity`}
        >
          Connect {provider}
        </button>
      )}
    </div>
  )
}
