// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: false },
  modules: [
    '@nuxtjs/tailwindcss',
    '@nuxtjs/google-fonts',
    'nuxt-icon'
  ],
  googleFonts: {
    families: {
      'Inter': [100, 200, 300, 400, 500, 600, 700, 800, 900],
      'Onest': [100, 200, 300, 400, 500, 600, 700]
    }
  },
  app: {
    pageTransition: { name: 'page', mode: 'out-in' },
    head: {
      title: 'Flizan - Full-Stack Developer Portfolio',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { 
          name: 'description', 
          content: 'Портфолио Full-Stack разработчика с 3-летним опытом в создании современных веб-приложений. Специализация в React, Vue.js, Node.js и других технологиях.'
        },
        { name: 'format-detection', content: 'telephone=no' },
        { name: 'robots', content: 'index, follow' },
        { name: 'author', content: 'Flizan' },
        { property: 'og:title', content: 'Flizan - Full-Stack Developer Portfolio' },
        { property: 'og:description', content: 'Портфолио Full-Stack разработчика с 3-летним опытом в создании современных веб-приложений.' },
        { property: 'og:type', content: 'website' },
        { name: 'twitter:card', content: 'summary_large_image' },
        { name: 'twitter:title', content: 'Flizan - Full-Stack Developer Portfolio' },
        { name: 'twitter:description', content: 'Портфолио Full-Stack разработчика с 3-летним опытом в создании современных веб-приложений.' }
      ],
      link: [
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
      ],
      script: [
        {
          innerHTML: `
            (function(m,e,t,r,i,k,a){m[i]=m[i]||function(){(m[i].a=m[i].a||[]).push(arguments)};
            m[i].l=1*new Date();
            for (var j = 0; j < document.scripts.length; j++) {if (document.scripts[j].src === r) { return; }}
            k=e.createElement(t),a=e.getElementsByTagName(t)[0],k.async=1,k.src=r,a.parentNode.insertBefore(k,a)})
            (window, document, "script", "https://mc.yandex.ru/metrika/tag.js", "ym");

            ym(99658046, "init", {
              clickmap:true,
              trackLinks:true,
              accurateTrackBounce:true,
              webvisor:true
            });
          `
        }
      ],
      noscript: [
        {
          innerHTML: '<div><img src="https://mc.yandex.ru/watch/99658046" style="position:absolute; left:-9999px;" alt="" /></div>'
        }
      ]
    }
  }
})