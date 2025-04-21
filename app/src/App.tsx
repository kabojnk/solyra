import { useState } from 'react'
import { usePWA } from './hooks/usePWA'
import './App.css'

const pills = [
  {
    label: "Hyperfocal",
    description: "Hyperfocal distance calculator",
    href: "#hyperfocal",
    color: "bg-primary text-white",
  },
  {
    label: "My Cameras & Lenses",
    description: "Track your camera bodies and lenses",
    href: "#gear",
    color: "bg-magenta text-white",
  },
  {
    label: "Weather Quality",
    description: "Check weather for photography",
    href: "#weather",
    color: "bg-yellow text-gray-900",
  },
  {
    label: "Composition Check",
    description: "Grid & overlay tool for composition",
    href: "#composition",
    color: "bg-primary-dark text-white",
  },
];

function App() {
  const [count, setCount] = useState(0)
  const { needRefresh, offlineReady } = usePWA()

  return (
    <div className="min-h-screen flex flex-col items-center justify-center px-2 py-6 bg-gradient-to-br from-primary via-magenta to-yellow">
      <h1 className="text-3xl sm:text-4xl font-bold mb-8 text-white drop-shadow-lg font-sans text-center">
        Solyra Photo Toolbox
      </h1>
      <div className="w-full max-w-sm flex flex-col gap-6">
        {pills.map((pill) => (
          <a
            key={pill.label}
            href={pill.href}
            className={`glass flex flex-col items-start p-5 ${pill.color} rounded-pill transition hover:scale-105 focus:outline-none focus:ring-4 focus:ring-white/40`}
            style={{backdropFilter: 'blur(12px)'}}
          >
            <span className="text-lg font-semibold font-sans">{pill.label}</span>
            <span className="text-sm opacity-80 font-sans">{pill.description}</span>
          </a>
        ))}
      </div>
      <footer className="mt-10 text-xs text-white/70 font-sans">
        &copy; {new Date().getFullYear()} Solyra
      </footer>
    </div>
  )
}

export default App
