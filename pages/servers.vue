<template>
    <div class="min-h-screen bg-[#0A0A0A] text-white">
      <div class="relative min-h-screen">
        <div class="absolute inset-0">
          <div class="absolute w-[800px] h-[800px] bg-blue-500/20 rounded-full filter blur-[120px] -top-96 -left-96 animate-pulse"></div>
          <div class="absolute w-[600px] h-[600px] bg-purple-500/20 rounded-full filter blur-[120px] top-1/2 -right-48 animate-pulse delay-1000"></div>
          <div class="absolute w-[500px] h-[500px] bg-emerald-500/20 rounded-full filter blur-[120px] -bottom-48 left-1/2 transform -translate-x-1/2 animate-pulse delay-2000"></div>
        </div>
  
        <div class="relative z-10 container mx-auto px-4 py-16">
          <div class="flex justify-between items-center mb-12">
            <div class="space-y-2">
              <h1 class="font-onest text-4xl font-bold bg-gradient-to-r from-blue-400 via-purple-400 to-emerald-400 bg-clip-text text-transparent">
                Мониторинг Серверов
              </h1>
              <p class="text-gray-400 text-sm">Отслеживание состояния и производительности серверов в реальном времени</p>
            </div>
            <NuxtLink 
              to="/" 
              class="nav-button group inline-flex items-center gap-2 px-6 py-3 bg-gray-900/50 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
            >
              <Icon name="mdi:arrow-left" class="w-5 h-5 transition-transform duration-300 group-hover:-translate-x-1" />
              Назад
            </NuxtLink>
          </div>
  
          <div v-if="loading && !hasInitialData" class="flex items-center justify-center py-20">
            <div class="relative">
              <div class="w-16 h-16 border-4 border-blue-400/20 border-t-blue-400 rounded-full animate-spin"></div>
              <div class="absolute inset-0 w-16 h-16 border-4 border-transparent border-t-purple-400 rounded-full animate-spin delay-150"></div>
            </div>
          </div>
  
          <div v-else-if="error && !hasInitialData" class="text-center py-20">
            <div class="bg-red-500/10 backdrop-blur-sm rounded-2xl p-8 max-w-md mx-auto border border-red-500/20">
              <Icon name="mdi:alert-circle" class="w-16 h-16 text-red-400 mx-auto mb-4" />
              <p class="text-lg text-gray-300 mb-4">{{ error }}</p>
              <button 
                @click="fetchServers"
                class="px-6 py-3 bg-red-500/20 rounded-xl hover:bg-red-500/30 transition-colors text-red-300 hover:text-red-200"
              >
                Попробовать снова
              </button>
            </div>
          </div>
  
          <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
            <div v-for="server in serversWithLocations" :key="server.ip" 
                 class="server-card group">
              <div class="relative bg-gray-900/50 backdrop-blur-sm rounded-xl p-6 border border-gray-800/50 transition-all duration-300">
                <div class="absolute inset-0 bg-gradient-to-br from-blue-500/0 to-purple-500/0 opacity-0 group-hover:opacity-10 rounded-xl transition-opacity duration-300"></div>
                
                <div class="flex items-center justify-between mb-6">
                  <div class="flex items-center gap-3">
                    <div class="relative">
                      <div class="w-3 h-3 rounded-full" 
                           :class="server.status === 'online' ? 'bg-emerald-400' : 'bg-red-400'"></div>
                      <div class="absolute inset-0 w-3 h-3 rounded-full animate-ping" 
                           :class="server.status === 'online' ? 'bg-emerald-400/50' : 'bg-red-400/50'"></div>
                    </div>
                    <h3 class="font-semibold text-lg group-hover:text-white transition-colors">{{ server.name }}</h3>
                  </div>
                  <span :class="[
                    'px-3 py-1 text-xs font-medium rounded-full transition-colors duration-300',
                    server.status === 'online' ? 'bg-emerald-500/20 text-emerald-400 group-hover:bg-emerald-500/30' : 'bg-red-500/20 text-red-400 group-hover:bg-red-500/30'
                  ]">
                    {{ server.status }}
                  </span>
                </div>
                
                <div class="space-y-4">
                  <div class="flex items-center justify-between p-3 rounded-lg bg-gray-800/50 group-hover:bg-gray-800/70 transition-colors">
                    <span class="text-gray-400">Локация:</span>
                    <div class="flex items-center gap-2">
                      <Icon :name="server.location.flag" class="w-5 h-5" />
                      <span class="text-gray-200 group-hover:text-white transition-colors">{{ server.location.name }}</span>
                    </div>
                  </div>
                  
                  <div class="flex items-center justify-between p-3 rounded-lg bg-gray-800/50 group-hover:bg-gray-800/70 transition-colors">
                    <span class="text-gray-400">Пинг:</span>
                    <span :class="[
                      'px-2 py-1 rounded-full text-xs font-medium transition-colors duration-300',
                      server.status === 'online' ? getPingClass(server.ping) : 'bg-gray-500/20 text-gray-400'
                    ]">
                      {{ server.status === 'online' ? `${server.ping.toFixed(1)} мс` : '?' }}
                    </span>
                  </div>
                </div>
  
                <div class="mt-6 pt-4 border-t border-gray-800/50">
                  <div class="flex items-center justify-between text-sm">
                    <span class="text-gray-400">IP:</span>
                    <span class="text-gray-300 font-mono text-sm">{{ server.ip }}</span>
                  </div>
                </div>
              </div>
            </div>
          </div>
  
          <div v-if="hasInitialData" class="flex justify-center mt-12">
            <button 
              @click="fetchServers"
              class="nav-button group inline-flex items-center gap-2 px-8 py-4 bg-gradient-to-r from-blue-500/20 to-purple-500/20 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
              :disabled="loading"
            >
              <Icon name="mdi:refresh" 
                    class="w-5 h-5 transition-transform duration-300"
                    :class="{ 'animate-spin': loading }" />
              Обновить статус
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  const API_URL = 'https://flizan.ru/api'
  
  const servers = ref([])
  const loading = ref(true)
  const error = ref(null)
  const hasInitialData = ref(false)
  
  const serverMap = {
    '46.174.48.244': {
      name: 'MSK-1',
      location: {
        name: 'Москва',
        flag: 'flag:ru-4x3',
        country: 'Россия'
      }
    },
    '46.174.49.153': {
      name: 'MSK-2',
      location: {
        name: 'Москва',
        flag: 'flag:ru-4x3',
        country: 'Россия'
      }
    },
    '89.39.121.159': {
      name: 'GER-1',
      location: {
        name: 'Франкфурт',
        flag: 'flag:de-4x3',
        country: 'Германия'
      }
    },
    '87.120.166.96': {
      name: 'GER-2',
      location: {
        name: 'Франкфурт',
        flag: 'flag:de-4x3',
        country: 'Германия'
      }
    }
  }
  
  const serversWithLocations = computed(() => {
    return servers.value.map(server => ({
      ...server,
      name: serverMap[server.ip]?.name || 'Неизвестный сервер',
      location: serverMap[server.ip]?.location || {
        name: 'Неизвестно',
        flag: 'mdi:earth',
        country: 'Неизвестно'
      }
    }))
  })
  
  const getPingClass = (ping) => {
    if (ping < 50) return 'bg-emerald-500/20 text-emerald-400 group-hover:bg-emerald-500/30'
    if (ping < 100) return 'bg-yellow-500/20 text-yellow-400 group-hover:bg-yellow-500/30'
    return 'bg-red-500/20 text-red-400 group-hover:bg-red-500/30'
  }
  
  const fetchServers = async () => {
    loading.value = true
    error.value = null
    
    try {
      const response = await fetch(`${API_URL}/status`)
      if (!response.ok) {
        throw new Error('Ошибка при получении данных с сервера')
      }
      servers.value = await response.json()
      hasInitialData.value = true
    } catch (err) {
      error.value = 'Не удалось загрузить информацию о серверах. Пожалуйста, попробуйте позже.'
      console.error('Error fetching servers:', err)
    } finally {
      loading.value = false
    }
  }
  
  onMounted(() => {
    fetchServers()
  })
  
  let pollInterval
  onMounted(() => {
    pollInterval = setInterval(fetchServers, 60000)
  })
  
  onUnmounted(() => {
    if (pollInterval) {
      clearInterval(pollInterval)
    }
  })
  </script>
  
  <style scoped>
  .server-card {
    transform: translateZ(0);
  }
  
  .server-card:hover > div {
    transform: translateY(-4px);
  }
  
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
  
  .delay-2000 {
    animation-delay: 2s;
  }
  </style>