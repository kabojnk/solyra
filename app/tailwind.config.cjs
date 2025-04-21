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
          DEFAULT: "#a259ff", // purple
          dark: "#6d28d9",
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
