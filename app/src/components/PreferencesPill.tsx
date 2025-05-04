import { usePreferences } from '../stores/preferences'

export function PreferencesPill() {
  const { nightMode, toggleNightMode } = usePreferences()

  return (
    <div className="flex items-center gap-2 rounded-full bg-surface px-4 py-2 backdrop-blur-sm">
      <label className="text-sm text-content nightmode:text-content-night">Night Mode</label>
      <button
        onClick={toggleNightMode}
        className="relative inline-flex h-6 w-11 items-center rounded-full transition-colors bg-gray-600 nightmode:bg-primary-accent-night"
      >
        <span
          className={`inline-block h-4 w-4 transform rounded-full bg-white transition-transform ${
            nightMode ? 'translate-x-6' : 'translate-x-1'
          }`}
        />
      </button>
    </div>
  )
}
