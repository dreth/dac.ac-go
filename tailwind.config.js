/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./generated/*.html",
  ],
  theme: {
    extend: {
      colors: {
        'links': '#0ff',
        'blog-links': 'rgb(0, 174, 255)',
        'links-hover': '#7dff7d',
        'blog-links-hover': '#62ff62',
        'background': 'black',
        'text': 'rgb(205, 205, 205)',
        'light': 'white',
        'text-light': 'black',
        'almost-black': 'rgba(7, 7, 7)',
        'h1-h3-article-color': 'rgb(129, 193, 253)',
        'h2-article-color': 'rgb(97, 115, 255)',
        'code-block': 'rgb(9, 8, 13)',
        'quote-block': 'rgb(39, 38, 43)',
        'playlists-mobile': '#9ce5f7',
      },
      fontSize: {
        'xxs': [
          '.6rem', {
            lineHeight: '0.7rem',
          }
        ],
      },
      screens: { 
        'xs': '374px',
      },
    }
  },
  plugins: [
    function({ addComponents }) {
      addComponents({
        '.custom-li-marker': {
          listStyleType: 'none',
          '&::before': {
            content: "'🔗 '",
            // Add other styles if needed
          },
        },
      })
    },
  ],
}
