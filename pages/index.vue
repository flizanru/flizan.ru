<template>
  <div class="min-h-screen bg-[#0A0A0A] text-white">
    <div class="relative min-h-screen flex items-center justify-center">
      <div class="absolute inset-0 bg-gradient-to-b from-purple-900/10 to-[#0A0A0A]"></div>
      <div class="relative z-10 text-center max-w-4xl mx-auto px-4">
        <h1 class="font-onest text-7xl font-bold mb-6 bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text text-transparent">
          Flizan
        </h1>
        <p class="font-inter text-2xl text-gray-300 mb-8">
          Full-Stack Разработчик
        </p>
        <p class="font-inter text-lg text-gray-400 max-w-2xl mx-auto leading-relaxed mb-12">
          С 3-летним опытом создания современных веб-приложений и решения сложных технических задач.
          Увлечен созданием инновационных решений и обеспечением исключительного пользовательского опыта.
        </p>
        <div class="flex items-center justify-center gap-4 flex-wrap">
          <NuxtLink 
            to="/services" 
            class="nav-button group inline-flex items-center gap-2 px-8 py-3 bg-gray-900/50 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
          >
            Услуги
            <Icon name="mdi:application-cog" class="w-5 h-5 transition-transform duration-300 group-hover:translate-x-1" />
          </NuxtLink>
          <NuxtLink 
            to="/skills" 
            class="nav-button group inline-flex items-center gap-2 px-8 py-3 bg-gray-900/50 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
          >
            Навыки
            <Icon name="mdi:arrow-right" class="w-5 h-5 transition-transform duration-300 group-hover:translate-x-1" />
          </NuxtLink>
          <NuxtLink 
            to="/projects" 
            class="nav-button group inline-flex items-center gap-2 px-8 py-3 bg-gray-900/50 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
          >
            Проекты
            <Icon name="mdi:code-braces" class="w-5 h-5 transition-transform duration-300 group-hover:translate-x-1" />
          </NuxtLink>
          <!-- <NuxtLink 
            to="/servers" 
            class="nav-button group inline-flex items-center gap-2 px-8 py-3 bg-gray-900/50 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
          >
            Сервера
            <Icon name="mdi:server" class="w-5 h-5 transition-transform duration-300 group-hover:translate-x-1" />
          </NuxtLink> -->
          <a 
            href="https://t.me/flizan_s" 
            target="_blank" 
            rel="noopener noreferrer"
            class="nav-button group inline-flex items-center gap-2 px-8 py-3 bg-blue-600/80 backdrop-blur-sm rounded-xl font-medium text-white hover:bg-blue-700/80 transition-all duration-300"
          >
            Связаться
            <Icon name="mdi:telegram" class="w-5 h-5 transition-transform duration-300 group-hover:translate-x-1" />
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.nav-button {
  box-shadow: inset 0 1px 0 0 rgba(255, 255, 255, 0.1);
}

.nav-button:hover {
  transform: translateY(-2px);
  box-shadow: inset 0 1px 0 0 rgba(255, 255, 255, 0.1),
              0 4px 20px rgba(0, 0, 0, 0.3);
}
</style>