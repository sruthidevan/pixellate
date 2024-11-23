<template>
    <div class="min-h-screen flex flex-col bg-gradient-to-b from-yellow-100 via-yellow-200 to-yellow-300">
      <div class="flex-grow flex flex-col items-center justify-center text-center p-8">
        <h1 class="text-4xl font-bold mb-4 text-black">About Us</h1>
        <p class="text-lg mb-6 text-gray-700 max-w-3xl">
          After years of crafting code and solving tough tech puzzles, I launched <span class="highlight">Pixellate</span>—a software development business committed to delivering clean code, tailored solutions, and personalized attention for every project.
        </p>
        <p class="text-lg text-gray-700 max-w-3xl">
          Whether you need custom software, cloud integrations, or backend systems, <span class="highlight">Pixellate</span> is here to help transform your vision into reality. Let’s build something amazing together.
        </p>
      </div>
    </div>
  </template>
  
  <script setup>
    // No custom logic needed for this static page.
  </script>
  
  <style scoped>
  /* Highlight class */
  .highlight {
    background-color: #ffcc99; /* Orangy/peachish highlight color */
    padding: 0.2rem 0.4rem; /* Add some padding */
    border-radius: 0.2rem; /* Slightly round the corners */
  }
  </style>