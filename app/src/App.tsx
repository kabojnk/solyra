import { useState } from 'react'
import { usePWA } from './hooks/usePWA'
import './App.css'

function App() {
  const [count, setCount] = useState(0)
  const { needRefresh, offlineReady } = usePWA()

  return (
      <>
        <h1 className="text-3xl font-bold underline">
          Hello world!
        </h1>
      </>
  )
}

export default App
