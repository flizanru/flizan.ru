<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>
</template>

<style>
body {
  margin: 0;
  padding: 0;
  background: #0A0A0A;
  overflow-y: auto;
}

.page-enter-active,
.page-leave-active {
  transition: all 0.5s cubic-bezier(0.4, 0, 0.2, 1);
}

.page-enter-from {
  opacity: 0;
  transform: scale(0.98);
}

.page-leave-to {
  opacity: 0;
  transform: scale(1.02);
}

/* Скрываем полосы прокрутки, но сохраняем функционал */
::-webkit-scrollbar {
  width: 0;
  background: transparent;
}

* {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>