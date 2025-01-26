<template>
  <div class="min-h-screen bg-[#0A0A0A] text-white">
    <div class="relative min-h-screen">
      <!-- Background Effects -->
      <div class="absolute inset-0 opacity-10">
        <div class="absolute w-[500px] h-[500px] bg-blue-500 rounded-full filter blur-[128px] -top-48 -left-24"></div>
        <div class="absolute w-[400px] h-[400px] bg-purple-500 rounded-full filter blur-[128px] top-1/2 -right-24"></div>
      </div>

      <!-- Content -->
      <div class="relative z-10 container mx-auto px-4 py-16">
        <div class="flex justify-between items-center mb-12">
          <h1 class="font-onest text-4xl font-bold bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">
            Мои Проекты
          </h1>
          <NuxtLink 
            to="/" 
            class="nav-button group inline-flex items-center gap-2 px-6 py-3 bg-gray-900/50 backdrop-blur-sm rounded-xl font-medium text-gray-300 hover:text-white transition-all duration-300"
          >
            <Icon name="mdi:arrow-left" class="w-5 h-5 transition-transform duration-300 group-hover:-translate-x-1" />
            Назад
          </NuxtLink>
        </div>

        <!-- Projects Grid -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="project in projects" :key="project.id" 
               class="project-card bg-gray-900/50 backdrop-blur-sm rounded-xl overflow-hidden border border-gray-800 transition-all duration-300 hover:border-transparent">
            <!-- Project Image -->
            <div class="relative aspect-video overflow-hidden">
              <img 
                :src="project.image" 
                :alt="project.title"
                class="w-full h-full object-cover transition-transform duration-500 hover:scale-110"
              />
            </div>

            <!-- Project Content -->
            <div class="p-6">
              <div class="flex items-start justify-between mb-4">
                <h3 class="font-onest text-xl font-semibold text-white">
                  {{ project.title }}
                </h3>
                <span 
                  class="px-3 py-1 text-xs font-medium rounded-full"
                  :class="getLanguageClass(project.language)"
                >
                  {{ project.language }}
                </span>
              </div>
              <p class="text-gray-400 text-sm mb-4">{{ project.description }}</p>
              
              <!-- Project Links -->
              <div class="flex gap-3">
                <a v-if="project.demoUrl" 
                   :href="project.demoUrl" 
                   target="_blank"
                   class="inline-flex items-center gap-1 text-sm text-blue-400 hover:text-blue-300 transition-colors"
                >
                  <Icon name="mdi:link-variant" class="w-4 h-4" />
                  Site
                </a>
                <a v-if="project.githubUrl" 
                   :href="project.githubUrl" 
                   target="_blank"
                   class="inline-flex items-center gap-1 text-sm text-gray-400 hover:text-white transition-colors"
                >
                  <Icon name="mdi:github" class="w-4 h-4" />
                  GitHub
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
const projects = [
  {
    id: 1,
    title: "Flizan Portfolio",
    language: "Vue.js",
    description: "Современное портфолио с анимациями и адаптивным дизайном. Использует Nuxt 3, Tailwind CSS и различные интерактивные элементы.",
    image: "http://flizan.ru/images/image.png",
    demoUrl: "https://flizan.ru",
    githubUrl: "https://github.com/flizanru/flizan.ru"
  },
  {
    id: 2,
    title: "Паспорт GambitRP",
    language: "Lua",
    description: "Польнофункциональный код системы паспорта в игре GarrysMod на сервере GambitRP",
    image: "http://flizan.ru/images/image2.png",
    demoUrl: "https://gambitrp.ru",
    // githubUrl: "https://github.com/example/ecommerce"
  },
   {
     id: 3,
     title: "Амням",
     language: "Амням",
     description: "Мне просто для красоты захотелось добавить третью стопку и я добавил сюда амняма",
     image: "https://pic.rutubelist.ru/video/46/6d/466d15d2f16f6e59150edcd44f5d88a3.jpg",
     demoUrl: "https://www.youtube.com/watch?v=sNLVxifWp_4",
     githubUrl: "https://www.youtube.com/watch?v=sNLVxifWp_4"
   }
]

const getLanguageClass = (language) => {
  const classes = {
    'Vue.js': 'bg-emerald-500/20 text-emerald-400',
    'React': 'bg-blue-500/20 text-blue-400',
    'Lua': 'bg-emerald-500/20 text-emerald-400',
    'Node.js': 'bg-green-500/20 text-green-400',
    'Python': 'bg-yellow-500/20 text-yellow-400',
    'Амням': 'bg-blue-500/20 text-blue-400',
    'JavaScript': 'bg-yellow-500/20 text-yellow-400'
  }
  return classes[language] || 'bg-gray-500/20 text-gray-400'
}
</script>

<style scoped>
.project-card {
  position: relative;
}

.project-card::before {
  content: '';
  position: absolute;
  inset: -1px;
  border-radius: 0.75rem;
  background: linear-gradient(
    45deg,
    rgba(59, 130, 246, 0.5),
    rgba(147, 51, 234, 0.5)
  );
  opacity: 0;
  transition: opacity 0.3s ease;
  pointer-events: none;
  z-index: -1;
}

.project-card:hover::before {
  opacity: 0.15;
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