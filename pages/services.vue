<template>
    <div class="min-h-screen bg-[#0A0A0A] text-white">
      <div class="relative min-h-screen">
        <div class="absolute inset-0">
          <div class="absolute w-[800px] h-[800px] bg-purple-500/20 rounded-full filter blur-[120px] -top-96 -left-96 animate-pulse"></div>
          <div class="absolute w-[600px] h-[600px] bg-blue-500/20 rounded-full filter blur-[120px] top-1/2 -right-48 animate-pulse delay-1000"></div>
          <div class="absolute w-[500px] h-[500px] bg-emerald-500/20 rounded-full filter blur-[120px] bottom-0 left-1/2 transform -translate-x-1/2 animate-pulse delay-2000"></div>
        </div>
  
        <div class="relative z-10 container mx-auto px-4 py-16">
          <div class="flex justify-between items-center mb-16">
            <h1 class="font-onest text-4xl font-bold bg-gradient-to-r from-purple-400 via-blue-400 to-emerald-400 bg-clip-text text-transparent">
              Услуги
            </h1>
            <NuxtLink 
              to="/" 
              class="nav-button group inline-flex items-center gap-2 px-6 py-3 bg-gray-900/50 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
            >
              <Icon name="mdi:arrow-left" class="w-5 h-5 transition-transform duration-300 group-hover:-translate-x-1" />
              Назад
            </NuxtLink>
          </div>
  
          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div v-for="service in services" :key="service.id" 
                 class="service-card group bg-gray-900/50 backdrop-blur-sm rounded-xl p-8 border border-gray-800 transition-all duration-300">
              <div class="flex items-center gap-4 mb-6">
                <div class="w-12 h-12 rounded-xl bg-gray-800/50 flex items-center justify-center">
                  <Icon :name="service.icon" class="w-6 h-6" :class="service.iconColor" />
                </div>
                <h3 class="text-xl font-semibold text-white">{{ service.title }}</h3>
              </div>
              
              <p class="text-gray-400 mb-6">{{ service.description }}</p>
              
              <div class="space-y-3 mb-6">
                <div v-for="(feature, index) in service.features" :key="index"
                     class="flex items-start gap-3">
                  <Icon name="mdi:check-circle" class="w-5 h-5 text-emerald-400 mt-0.5" />
                  <span class="text-gray-300 text-sm">{{ feature }}</span>
                </div>
              </div>
  
              <div class="pt-6 border-t border-gray-800">
                <div class="flex items-end gap-2 mb-2">
                  <span class="text-2xl font-bold text-white">от {{ service.price.from.toLocaleString() }} ₽</span>
                  <span v-if="service.price.to" class="text-gray-400">
                    до {{ service.price.to.toLocaleString() }} ₽
                  </span>
                </div>
                <p class="text-sm text-gray-500">{{ service.priceNote }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  const services = [
    {
      id: 1,
      title: 'Разработка веб-сайтов',
      icon: 'mdi:web',
      iconColor: 'text-blue-400',
      description: 'Создание современных, адаптивных и интуитивно понятных веб-сайтов с полным циклом разработки.',
      features: [
        'Frontend разработка (HTML, CSS, JavaScript, современные фреймворки)',
        'Backend разработка (Go, PHP)',
        'Интеграция баз данных и API',
        'Адаптивный дизайн для всех устройств'
      ],
      price: {
        from: 5000,
        to: 50000
      },
      priceNote: 'Стоимость зависит от сложности проекта и требуемого функционала'
    },
    {
      id: 2,
      title: 'Администрирование серверов',
      icon: 'mdi:server',
      iconColor: 'text-purple-400',
      description: 'Профессиональная настройка и управление серверной инфраструктурой.',
      features: [
        'Настройка панелей управления (Pterodactyl)',
        'Установка и настройка Linux серверов',
        'Оптимизация производительности',
        'Мониторинг и обеспечение безопасности'
      ],
      price: {
        from: 1000,
        to: 5000
      },
      priceNote: 'Ежемесячное обслуживание от 5000 ₽'
    },
    {
      id: 3,
      title: 'Разработка ботов',
      icon: 'mdi:robot',
      iconColor: 'text-emerald-400',
      description: 'Создание и настройка ботов для Discord и Telegram с расширенным функционалом.',
      features: [
        'Интеграция с различными API',
        'Автоматизация задач',
        'Кастомные команды',
        'Панель управления ботом'
      ],
      price: {
        from: 1000,
        to: 10000
      },
      priceNote: 'Техническая поддержка и обновления включены'
    },
    {
      id: 4,
      title: 'Разработка на Lua',
      icon: 'mdi:script-text',
      iconColor: 'text-yellow-400',
      description: 'Профессиональная разработка скриптов и модификаций на языке Lua.',
      features: [
        'Написание скриптов для игровых платформ',
        'Создание модификаций',
        'Автоматизация процессов',
        'Оптимизация производительности'
      ],
      price: {
        from: 5000,
        to: 30000
      },
      priceNote: 'Цена зависит от сложности и объема работ'
    }
  ]
  </script>
  
  <style scoped>
  .service-card {
    position: relative;
    overflow: hidden;
  }
  
  .service-card::before {
    content: '';
    position: absolute;
    inset: -1px;
    background: linear-gradient(
      to right bottom,
      rgba(139, 92, 246, 0.1),
      rgba(59, 130, 246, 0.1)
    );
    border-radius: 0.75rem;
    opacity: 0;
    transition: opacity 0.3s ease;
  }
  
  .service-card:hover::before {
    opacity: 1;
  }
  
  .service-card:hover {
    transform: translateY(-4px);
    border-color: transparent;
    box-shadow: 
      0 4px 24px -1px rgba(0, 0, 0, 0.2),
      0 0 0 1px rgba(139, 92, 246, 0.1);
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