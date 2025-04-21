/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    fontFamily: {
      sans: ["Inter", "ui-sans-serif", "system-ui", "sans-serif"],
    },
    extend: {
      colors: {
        primary: {
          DEFAULT: "#6f42c1", // deep purple
          dark: "#2d1a4d", // midnight purple
        },
        midnight: {
          DEFAULT: "#191933", // deep midnight blue
        },
        magenta: {
          DEFAULT: "#ff3cac",
        },
        yellow: {
          DEFAULT: "#ffe156",
        },
        glass: {
          DEFAULT: "rgba(255,255,255,0.15)",
        },
      },
      backdropBlur: {
        xs: '2px',
      },
      borderRadius: {
        pill: '9999px',
      },
    },
  },
  plugins: [],
};
