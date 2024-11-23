<template>
    <div class="min-h-screen flex items-center justify-center bg-gradient-to-b from-yellow-100 via-yellow-200 to-yellow-300">
      <div class="text-center p-8">
        <h1 class="text-5xl font-bold mb-4 text-black">Welcome to Pixellate</h1>
        <p class="text-lg mb-6 text-gray-700">
            Turning your tech dreams into clean, scalable, and efficient software solutions. From backend systems to cloud integrations, we bring your ideas to life.
        </p>
        <div class="flex justify-center space-x-4">
          <router-link to="/services" class="btn">Explore Services</router-link>
          <router-link to="/contact" class="btn">Contact Us</router-link>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  // No custom logic needed for this static page.
  </script>
  
  <style scoped>
  /* Button styles */
  .btn {
    padding: 10px 20px;
    background: linear-gradient(to right, #ffcc00, #ff9900);
    color: white;
    font-weight: bold;
    border-radius: 25px;
    text-transform: uppercase;
    transition: transform 0.2s, box-shadow 0.2s;
    text-decoration: none;
  }
  
  .btn:hover {
    transform: scale(1.05);
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
  }
  </style>
  