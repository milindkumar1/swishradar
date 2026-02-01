export default function Home() {
  return (
    <main className="min-h-screen bg-gradient-to-b from-gray-900 to-gray-800">
      <div className="container mx-auto px-4 py-16">
        {/* Hero Section */}
        <div className="text-center mb-16">
          <h1 className="text-6xl font-bold text-white mb-4">
            🏀 SwishRadar
          </h1>
          <p className="text-xl text-gray-300 mb-8">
            Advanced Analytics for ESPN Fantasy Basketball
          </p>
          <div className="flex justify-center gap-4">
            <a
              href="/dashboard"
              className="bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-8 rounded-lg transition-colors"
            >
              Get Started
            </a>
            <a
              href="/docs"
              className="bg-gray-700 hover:bg-gray-600 text-white font-bold py-3 px-8 rounded-lg transition-colors"
            >
              Documentation
            </a>
          </div>
        </div>

        {/* Features Grid */}
        <div className="grid md:grid-cols-2 lg:grid-cols-4 gap-8 mb-16">
          <FeatureCard
            title="Trade Calculator"
            description="Evaluate trade fairness with advanced multi-category analysis"
            icon="🔄"
          />
          <FeatureCard
            title="Waiver Wire"
            description="AI-powered streaming recommendations based on games and trends"
            icon="📊"
          />
          <FeatureCard
            title="Power Rankings"
            description="Monte Carlo simulations for playoff odds and predictions"
            icon="🏆"
          />
          <FeatureCard
            title="Backtesting"
            description="Validate algorithm accuracy with historical data"
            icon="📈"
          />
        </div>

        {/* Quick Stats */}
        <div className="bg-gray-800 rounded-lg p-8 text-center">
          <h2 className="text-2xl font-bold text-white mb-6">
            Coming Soon
          </h2>
          <p className="text-gray-300">
            SwishRadar is currently in development. Check back soon for the full experience!
          </p>
        </div>
      </div>
    </main>
  );
}

function FeatureCard({ title, description, icon }: { title: string; description: string; icon: string }) {
  return (
    <div className="bg-gray-800 rounded-lg p-6 hover:bg-gray-750 transition-colors">
      <div className="text-4xl mb-4">{icon}</div>
      <h3 className="text-xl font-bold text-white mb-2">{title}</h3>
      <p className="text-gray-400">{description}</p>
    </div>
  );
}
