import { ReactNode, useEffect } from 'react'
import { usePreferences } from '../stores/preferences'

interface ThemeProviderProps {
  children: ReactNode
}

export function ThemeProvider({ children }: ThemeProviderProps) {
  const { nightMode } = usePreferences()

  useEffect(() => {
    if (nightMode) {
      document.documentElement.dataset.theme = 'nightmode'
    } else {
      document.documentElement.dataset.theme = ''
    }
  }, [nightMode])

  return <>{children}</>
}
