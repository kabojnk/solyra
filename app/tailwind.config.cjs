/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  darkMode: 'class',
  theme: {
    fontFamily: {
      sans: ["Inter", "ui-sans-serif", "system-ui", "sans-serif"],
    },
    extend: {
      colors: {
        background: {
          DEFAULT: "#191933", // deep midnight blue
          secondary: "#2d1a4d", // midnight purple
        },
        primary: {
          DEFAULT: "#6f42c1", // deep purple
          accent: "#ff3cac", // magenta
          highlight: "#ffe156", // yellow
        },
        primaryAccent: {
          DEFAULT: "#ff3cac", // magenta
        },
        surface: {
          DEFAULT: "rgba(255,255,255,0.15)", // glass effect
        },
      },
      backdropBlur: {
        xs: '2px',
      },
    },
  },
  plugins: [],
}
