/** @type {import('tailwindcss').Config} */
import daisyui from "daisyui";
export default {
  content: ["./src/**/*.{ts,tsx,js,jsx,html}"],
  theme: {
    extend: {},
  },
  plugins: [daisyui],
  daisyui:{
    themes:["emerald"]
  }
}

