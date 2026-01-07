import Link from 'next/link'
import { Music, ArrowRight, Check } from 'lucide-react'

export default function Home() {
  return (
    <main className="min-h-screen bg-gradient-to-br from-purple-900 via-blue-900 to-black">
      {/* Hero Section */}
      <div className="container mx-auto px-4 py-16">
        <div className="text-center mb-16">
          <div className="flex justify-center mb-6">
            <Music className="w-20 h-20 text-purple-400" />
          </div>
          <h1 className="text-6xl font-bold text-white mb-6">
            Music Migration Tool
          </h1>
          <p className="text-2xl text-gray-300 mb-8">
            Transfer your playlists between Spotify and Apple Music seamlessly
          </p>
          <Link
            href="/dashboard"
            className="inline-flex items-center px-8 py-4 bg-purple-600 hover:bg-purple-700 text-white text-lg font-semibold rounded-lg transition-colors"
          >
            Get Started
            <ArrowRight className="ml-2 w-5 h-5" />
          </Link>
        </div>

        {/* Features Grid */}
        <div className="grid md:grid-cols-3 gap-8 mt-20">
          <FeatureCard
            title="Fast Migration"
            description="Transfer hundreds of songs in minutes with our parallel processing engine"
            icon="⚡"
          />
          <FeatureCard
            title="Smart Matching"
            description="Advanced algorithm matches songs accurately using ISRC and metadata"
            icon="🎯"
          />
          <FeatureCard
            title="Real-time Progress"
            description="Watch your migration progress live with instant status updates"
            icon="📊"
          />
        </div>

        {/* How It Works */}
        <div className="mt-20">
          <h2 className="text-4xl font-bold text-white text-center mb-12">
            How It Works
          </h2>
          <div className="grid md:grid-cols-4 gap-6">
            <Step number={1} title="Connect" description="Authorize both accounts" />
            <Step number={2} title="Select" description="Choose playlists to migrate" />
            <Step number={3} title="Migrate" description="Let us do the work" />
            <Step number={4} title="Enjoy" description="Your music is ready!" />
          </div>
        </div>

        {/* Supported Platforms */}
        <div className="mt-20 text-center">
          <h2 className="text-3xl font-bold text-white mb-8">
            Supported Platforms
          </h2>
          <div className="flex justify-center items-center gap-12">
            <div className="bg-green-500 rounded-full p-6">
              <span className="text-4xl font-bold text-white">Spotify</span>
            </div>
            <div className="text-gray-400 text-2xl">↔</div>
            <div className="bg-gradient-to-br from-pink-500 to-red-500 rounded-full p-6">
              <span className="text-4xl font-bold text-white">Apple Music</span>
            </div>
          </div>
        </div>
      </div>
    </main>
  )
}

function FeatureCard({ title, description, icon }: { title: string; description: string; icon: string }) {
  return (
    <div className="bg-white/10 backdrop-blur-sm rounded-xl p-6 border border-white/20">
      <div className="text-5xl mb-4">{icon}</div>
      <h3 className="text-xl font-semibold text-white mb-2">{title}</h3>
      <p className="text-gray-300">{description}</p>
    </div>
  )
}

function Step({ number, title, description }: { number: number; title: string; description: string }) {
  return (
    <div className="text-center">
      <div className="w-16 h-16 bg-purple-600 rounded-full flex items-center justify-center mx-auto mb-4">
        <span className="text-2xl font-bold text-white">{number}</span>
      </div>
      <h3 className="text-lg font-semibold text-white mb-2">{title}</h3>
      <p className="text-gray-400 text-sm">{description}</p>
    </div>
  )
}
