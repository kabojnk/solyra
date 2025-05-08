import React from 'react'
import { BrowserRouter, Link, Route, Routes } from 'react-router-dom'
import { usePreferences } from './stores/preferences.ts'
import { PreferencesPill } from './components/PreferencesPill'
import { ThemeProvider } from './components/ThemeProvider'
import { usePWA } from './hooks/usePWA'
import Composition from './pages/Composition.tsx'
import Gear from './pages/Gear.tsx'
import Hyperfocal from './pages/Hyperfocal.tsx'
import NotFound from './pages/NotFound.tsx'
import Weather from './pages/Weather.tsx'
import Starfield from './Starfield'

const pills = [
  {
    label: 'Hyperfocal',
    description: 'Hyperfocal distance calculator',
    to: '/hyperfocal',
  },
  {
    label: 'My Cameras & Lenses',
    description: 'Track your camera bodies and lenses',
    to: '/gear',
  },
  {
    label: 'Weather Quality',
    description: 'Check weather for photography',
    to: '/weather',
  },
  {
    label: 'Composition Check',
    description: 'Grid & overlay tool for composition',
    to: '/composition',
  },
]

function Home() {
  const { needRefresh, offlineReady } = usePWA()
  const { nightMode } = usePreferences()
  return (
      <div className="flex flex-col items-center justify-center min-h-screen w-screen bg-gradient-to-br from-background nightmode:from-[#200] nightmode:via-background-secondary-night to-primary nightmode:to-[#300]">
        {!nightMode &&
          <>
            <Starfield/>
            <div className="w-screen h-screen bg-[url(/bg-main.jpg)] bg-cover bg-center fixed opacity-50"/>
          </>
        }
        <div className="min-h-screen flex flex-col items-center justify-center px-2 py-6">
          <h1 className="text-3xl sm:text-4xl font-bold mb-8 text-content nightmode:text-content-night drop-shadow-lg font-sans text-center">
            Solyra Photo Toolbox
          </h1>
          <div className="w-full max-w-sm flex flex-col gap-6">
            {pills.map((pill) => (
                <Link
                    key={pill.to}
                    to={pill.to}
                    className="bg-background/50 nightmode:bg-background-night/80 border-1 border-transparent nightmode:border-primary-night/30 text-content nightmode:text-primary-night rounded-full px-6 py-4 hover:scale-105 backdrop-blur-sm"
                >
                  <h2 className="text-lg font-bold">{pill.label}</h2>
                  <p className="text-sm opacity-80">{pill.description}</p>
                </Link>
            ))}
          </div>
          <div className="mt-8">
            <PreferencesPill/>
          </div>
        </div>
      </div>
  )
}

function App() {
  return (
      <BrowserRouter>
        <ThemeProvider>
        <Routes>
            <Route path="/" element={<Home/>}/>
            <Route path="/hyperfocal" element={<Hyperfocal/>}/>
            <Route path="/gear" element={<Gear/>}/>
            <Route path="/weather" element={<Weather/>}/>
            <Route path="/composition" element={<Composition/>}/>
            <Route path="*" element={<NotFound/>}/>
          </Routes>
        </ThemeProvider>
      </BrowserRouter>
  )
}

export default App
