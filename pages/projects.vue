<template>
    <div class="min-h-screen bg-[#0A0A0A] text-white">
      <div class="relative min-h-screen flex flex-col">
        <div class="absolute inset-0 opacity-10">
          <div class="absolute w-[500px] h-[500px] bg-blue-500 rounded-full filter blur-[128px] -top-48 -left-24"></div>
          <div class="absolute w-[400px] h-[400px] bg-purple-500 rounded-full filter blur-[128px] top-1/2 -right-24"></div>
        </div>
  
        <div class="relative flex-grow flex items-center justify-center z-10 px-4">
          <div class="text-center">
            <div class="mb-8">
              <Icon name="mdi:tools" class="w-24 h-24 mx-auto text-blue-400 animate-pulse" />
            </div>
            <h1 class="font-onest text-4xl font-bold mb-6 bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">
              Страница в разработке
            </h1>
            <p class="text-gray-400 mb-8 max-w-md mx-auto">
              В настоящее время мы работаем над созданием этой страницы. Скоро здесь появится портфолио проектов.
            </p>
            <NuxtLink 
              to="/" 
              class="nav-button group inline-flex items-center gap-2 px-8 py-3 bg-gray-900/50 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
            >
              <Icon name="mdi:arrow-left" class="w-5 h-5 transition-transform duration-300 group-hover:-translate-x-1" />
              Вернуться на главную
            </NuxtLink>
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