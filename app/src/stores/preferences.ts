import { create } from 'zustand'
import { persist } from 'zustand/middleware'
import localforage from 'localforage'

interface PreferencesState {
  nightMode: boolean
  toggleNightMode: () => void
}

export const usePreferences = create<PreferencesState>()(
  persist(
    (set) => ({
      nightMode: false,
      toggleNightMode: () => set((state) => ({ nightMode: !state.nightMode })),
    }),
    {
      name: 'solyra-preferences',
      storage: {
        // @ts-ignore
        getItem: async (name): Promise<unknown> => {
          const value = await localforage.getItem(name) as PreferencesState
          return value ?? null
        },
        setItem: async (name, value) => {
          await localforage.setItem(name, value)
        },
        removeItem: async (name) => {
          await localforage.removeItem(name)
        },
      },
    }
  )
)
