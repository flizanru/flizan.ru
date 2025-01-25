<template>
    <div class="min-h-screen bg-[#0A0A0A] text-white relative overflow-hidden">
      <div class="absolute inset-0 opacity-10">
        <div class="absolute w-[500px] h-[500px] bg-purple-500 rounded-full filter blur-[128px] -top-48 -left-24"></div>
        <div class="absolute w-[400px] h-[400px] bg-blue-500 rounded-full filter blur-[128px] top-1/2 -right-24"></div>
        <!-- <div class="absolute w-[600px] h-[600px] bg-emerald-500 rounded-full filter blur-[128px] -bottom-48 left-1/2 transform -translate-x-1/2"></div> -->
      </div>
  
      <div class="relative py-16 px-4 z-10">
        <div class="max-w-6xl mx-auto">
          <div class="flex justify-between items-center mb-16">
            <h2 class="font-onest text-4xl font-bold text-center">Технические Навыки</h2>
            <NuxtLink 
              to="/" 
              class="nav-button group inline-flex items-center gap-2 px-6 py-3 bg-gray-900 rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
            >
              <Icon name="mdi:arrow-left" class="w-5 h-5 transition-transform duration-300 group-hover:-translate-x-1" />
              Назад
            </NuxtLink>
          </div>
          
          <div class="mb-16">
            <h3 class="font-onest text-2xl font-semibold mb-8 flex items-center justify-center">
              <Icon name="mdi:monitor-dashboard" class="mr-3 text-purple-400" size="28" />
              Frontend Разработка
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div v-for="skill in sortedFrontendSkills" 
                   :key="skill.name"
                   class="skill-card bg-gray-900/50 backdrop-blur-sm rounded-xl p-6 border border-gray-800 transition-all duration-300 group hover:border-transparent"
                   :style="{ '--glow-color': getGlowColor(skill.level) }"
                   >
                <div class="flex items-center mb-4">
                  <img :src="skill.icon" :alt="skill.name" class="w-8 h-8 mr-3">
                  <h4 class="font-semibold text-lg transition-colors" :style="{ color: getTextColor(skill.level) }">{{ skill.name }}</h4>
                </div>
                <div class="flex items-center gap-2 mb-4">
                  <span :class="[
                    'px-2 py-1 rounded-full text-xs font-medium',
                    getSkillLevelClass(skill.level)
                  ]">
                    {{ skill.level }}
                  </span>
                </div>
                <p class="text-gray-400 text-sm mb-4">{{ skill.brief }}</p>
                <p class="text-sm text-gray-300">{{ skill.description }}</p>
              </div>
            </div>
          </div>
  
          <div class="mb-16">
            <h3 class="font-onest text-2xl font-semibold mb-8 flex items-center justify-center">
              <Icon name="mdi:server" class="mr-3 text-blue-400" size="28" />
              Backend Разработка
            </h3>
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
              <div v-for="skill in sortedBackendSkills" 
                   :key="skill.name"
                   class="skill-card bg-gray-900/50 backdrop-blur-sm rounded-xl p-6 border border-gray-800 transition-all duration-300 group hover:border-transparent"
                   :style="{ '--glow-color': getGlowColor(skill.level) }"
                   >
                <div class="flex items-center mb-4">
                  <img :src="skill.icon" :alt="skill.name" class="w-8 h-8 mr-3">
                  <h4 class="font-semibold text-lg transition-colors" :style="{ color: getTextColor(skill.level) }">{{ skill.name }}</h4>
                </div>
                <div class="flex items-center gap-2 mb-4">
                  <span :class="[
                    'px-2 py-1 rounded-full text-xs font-medium',
                    getSkillLevelClass(skill.level)
                  ]">
                    {{ skill.level }}
                  </span>
                </div>
                <p class="text-gray-400 text-sm mb-4">{{ skill.brief }}</p>
                <p class="text-sm text-gray-300">{{ skill.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </div>
  
      <footer class="relative py-6 z-10 border-t border-gray-800/20">
        <div class="text-center">
          <p class="font-inter text-sm text-gray-500">
            © {{ new Date().getFullYear() }} Flizan. Все права защищены.
          </p>
        </div>
      </footer>
    </div>
  </template>
  
  <script setup>
  import skillsConfig from '../config/skills.json'
  
  const getSkillLevelClass = (level) => {
    const classes = {
      '5 лет': 'bg-emerald-500/20 text-emerald-400',
      '3 года': 'bg-violet-500/20 text-violet-400',
      '2 года': 'bg-blue-500/20 text-blue-400',
      '1 год': 'bg-indigo-500/20 text-indigo-400',
      '6 месяцев': 'bg-amber-500/20 text-amber-400',
      'Начинаю изучение': 'bg-gray-500/20 text-gray-400'
    }
    return classes[level] || 'bg-gray-500/20 text-gray-400'
  }
  
  const getGlowColor = (level) => {
    const colors = {
      '5 лет': 'rgb(16 185 129)', // emerald-500
      '3 года': 'rgb(139 92 246)', // violet-500
      '2 года': 'rgb(59 130 246)', // blue-500
      '1 год': 'rgb(99 102 241)',  // indigo-500
      '6 месяцев': 'rgb(245 158 11)', // amber-500
      'Начинаю изучение': 'rgb(107 114 128)' // gray-500
    }
    return colors[level] || 'rgb(31 41 55)' // gray-800
  }
  
  const getTextColor = (level) => {
    const colors = {
      '5 лет': 'rgb(52 211 153)', // emerald-400
      '3 года': 'rgb(167 139 250)', // violet-400
      '2 года': 'rgb(96 165 250)', // blue-400
      '1 год': 'rgb(129 140 248)', // indigo-400
      '6 месяцев': 'rgb(251 191 36)', // amber-400
      'Начинаю изучение': 'rgb(156 163 175)' // gray-400
    }
    return colors[level] || 'rgb(156 163 175)' // gray-400
  }
  
  const sortSkillsByExperience = (skills) => {
    const experienceOrder = {
      '5 лет': 5,
      '3 года': 4,
      '2 года': 3,
      '1 год': 2,
      '6 месяцев': 1,
      'Начинаю изучение': 0
    }
    
    return [...skills].sort((a, b) => experienceOrder[b.level] - experienceOrder[a.level])
  }
  
  const sortedFrontendSkills = sortSkillsByExperience(skillsConfig.frontend)
  const sortedBackendSkills = sortSkillsByExperience(skillsConfig.backend)
  </script>
  
  <style>
  .skill-card {
    position: relative;
  }
  
  .skill-card::before {
    content: '';
    position: absolute;
    inset: -1px;
    border-radius: 0.75rem;
    background: radial-gradient(
      800px circle at var(--mouse-x, 50%) var(--mouse-y, 50%),
      var(--glow-color),
      transparent 40%
    );
    opacity: 0;
    transition: opacity 0.3s ease;
    pointer-events: none;
    z-index: -1;
  }
  
  .skill-card:hover::before {
    opacity: 0.15;
  }
  
  .skill-card::after {
    content: '';
    position: absolute;
    inset: -1px;
    border-radius: 0.75rem;
    background: var(--glow-color);
    opacity: 0;
    transition: opacity 0.3s ease;
    pointer-events: none;
    filter: blur(16px);
    z-index: -2;
  }
  
  .skill-card:hover::after {
    opacity: 0.1;
  }
  
  .nav-button {
    box-shadow: inset 0 1px 0 0 rgba(255, 255, 255, 0.1);
  }
  
  .nav-button:hover {
    transform: translateY(-2px);
    box-shadow: inset 0 1px 0 0 rgba(255, 255, 255, 0.1),
                0 4px 20px rgba(0, 0, 0, 0.3);
  }
  </style>
  
  <script>
  if (process.client) {
    window.addEventListener('mousemove', (e) => {
      document.querySelectorAll('.skill-card').forEach(card => {
        const rect = card.getBoundingClientRect()
        const x = e.clientX - rect.left
        const y = e.clientY - rect.top
        
        card.style.setProperty('--mouse-x', `${x}px`)
        card.style.setProperty('--mouse-y', `${y}px`)
      })
    })
  }
  </script>