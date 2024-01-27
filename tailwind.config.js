/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.html"],
  theme: {
    extend: {
      spacing: {
        'content': '980px',
        'content-sm': '87.5%'
      },
      colors: {
        
      }
    },
  },
  plugins: [
    require("./plugin")({
      prefix: "flx",
      defaultFlavour: "flexoki_light",
    }),
    require("tailwind-fluid-typography")
  ],
}