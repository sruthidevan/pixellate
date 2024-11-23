<template>
    <div class="min-h-screen flex flex-col bg-gradient-to-b from-yellow-100 via-yellow-200 to-yellow-300">
      <div class="flex-grow p-8 text-center">
        <h1 class="text-4xl font-bold mb-6 text-black">Our Services</h1>
        <p class="text-lg mb-8 text-gray-700 max-w-2xl mx-auto">
          Explore the range of services Pixellate offers to bring your ideas to life.
        </p>
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <div
            v-for="service in services"
            :key="service.name"
            class="p-6 bg-white text-black rounded-lg shadow-md hover:shadow-xl transform hover:scale-105 transition-transform border-l-4 border-yellow-500"
          >
            <h2 class="text-2xl font-bold mb-4 text-yellow-600 underline">
              {{ service.name }}
            </h2>
            <p class="text-gray-700 text-lg">
              {{ service.description }}
            </p>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted } from 'vue';
  import axios from 'axios';
  
  const services = ref([]);
  
  onMounted(async () => {
    try {
      const response = await axios.get('http://localhost:8080/api/services'); // Ensure this matches your backend endpoint
      services.value = response.data;
    } catch (error) {
      console.error('Failed to fetch services:', error);
    }
  });
  </script>
  
  <style scoped>
  /* Styling for service cards */
  h2 {
    font-family: 'Montserrat', sans-serif;
    text-shadow: 1px 1px 3px rgba(0, 0, 0, 0.2); /* Add a subtle shadow for depth */
  }
  
  .grid > div {
    transition: all 0.3s ease-in-out; /* Smooth scaling animation */
  }
  
  .grid > div:hover {
    transform: translateY(-5px); /* Subtle hover effect */
    box-shadow: 0 10px 20px rgba(0, 0, 0, 0.1); /* Enhance hover shadow */
  }
  </style>