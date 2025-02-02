<template>
  <div class="min-h-screen bg-[#0A0A0A] text-white">
    <div class="relative min-h-screen">
      <!-- Background Effects -->
      <div class="absolute inset-0">
        <div class="absolute w-[800px] h-[800px] bg-purple-500/20 rounded-full filter blur-[120px] -top-96 -left-96 animate-pulse"></div>
        <div class="absolute w-[600px] h-[600px] bg-blue-500/20 rounded-full filter blur-[120px] top-1/2 -right-48 animate-pulse delay-1000"></div>
      </div>

      <!-- Content -->
      <div class="relative z-10 container mx-auto px-4">
        <div class="min-h-screen flex flex-col items-center justify-center">
          <!-- Header with Back Button -->
          <div class="absolute top-16 left-4 right-4">
            <div class="max-w-6xl mx-auto flex justify-between items-center">
              <h1 class="font-onest text-4xl font-bold bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text text-transparent">
                Проекты
              </h1>
              <NuxtLink 
                to="/" 
                class="nav-button group inline-flex items-center gap-2 px-6 py-3 bg-gray-900/50 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
              >
                <Icon name="mdi:arrow-left" class="w-5 h-5 transition-transform duration-300 group-hover:-translate-x-1" />
                Назад
              </NuxtLink>
            </div>
          </div>

          <!-- Maintenance Message -->
          <div class="text-center px-4 -mt-16">
            <div class="relative">
              <!-- Decorative Elements -->
              <div class="absolute inset-0 blur-3xl opacity-30">
                <div class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-48 h-48 bg-gradient-to-r from-purple-500 to-blue-500 rounded-full"></div>
              </div>

              <!-- Icon -->
              <div class="relative mb-8">
                <div class="w-24 h-24 mx-auto bg-gray-900/50 backdrop-blur-sm rounded-2xl flex items-center justify-center border border-gray-800">
                  <Icon 
                    name="mdi:tools" 
                    class="w-12 h-12 text-gray-400 animate-bounce"
                  />
                </div>
              </div>

              <!-- Message -->
              <h2 class="text-3xl font-bold mb-4 bg-gradient-to-r from-purple-400 to-blue-400 bg-clip-text text-transparent">
                Раздел в разработке
              </h2>
              <p class="text-gray-400 text-lg mb-8 max-w-md mx-auto">
                Мы работаем над созданием потрясающей галереи проектов. 
                Скоро здесь появится что-то интересное!
              </p>

              <!-- Status -->
              <div class="inline-flex items-center gap-2 px-4 py-2 bg-gray-900/50 backdrop-blur-sm rounded-full border border-gray-800">
                <div class="w-2 h-2 bg-blue-400 rounded-full animate-pulse"></div>
                <span class="text-sm text-gray-400">Ведутся работы</span>
              </div>
            </div>
          </div>
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

@keyframes pulse {
  0%, 100% { opacity: 0.2; }
  50% { opacity: 0.3; }
}

.animate-pulse {
  animation: pulse 4s cubic-bezier(0.4, 0, 0.6, 1) infinite;
}

.delay-1000 {
  animation-delay: 1s;
}
</style>