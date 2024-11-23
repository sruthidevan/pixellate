<template>
    <div class="min-h-screen flex items-center justify-center bg-gradient-to-b from-yellow-100 via-yellow-200 to-yellow-300">
      <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-lg">
        <h1 class="text-3xl font-bold mb-6 text-center text-black">Contact Us</h1>
        <form @submit.prevent="submitForm" class="space-y-6">
          <!-- Name Field -->
          <div class="flex items-center gap-4">
            <label for="name" class="block text-sm font-medium text-gray-700 w-36 text-left">
              Your Name
            </label>
            <input
              type="text"
              id="name"
              v-model="form.name"
              class="p-3 border border-gray-300 rounded focus:ring-2 focus:ring-yellow-400 focus:outline-none flex-1"
              placeholder="Enter your name"
              required
            />
          </div>
  
          <!-- Email Field -->
          <div class="flex items-center gap-4">
            <label for="email" class="block text-sm font-medium text-gray-700 w-36 text-left">
              Your Email
            </label>
            <input
              type="email"
              id="email"
              v-model="form.email"
              class="p-3 border border-gray-300 rounded focus:ring-2 focus:ring-yellow-400 focus:outline-none flex-1"
              placeholder="Enter your email"
              required
            />
          </div>
  
          <!-- Message Field -->
          <div class="flex items-center gap-4">
            <label for="message" class="block text-sm font-medium text-gray-700 w-36 text-left">
              Your Message
            </label>
            <textarea
              id="message"
              v-model="form.message"
              class="p-3 border border-gray-300 rounded focus:ring-2 focus:ring-yellow-400 focus:outline-none flex-1"
              rows="4"
              placeholder="Enter your message"
              required
            ></textarea>
          </div>
  
          <!-- Submit Button -->
          <button type="submit" class="btn w-full">Send Message</button>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue';
  
  const form = ref({
    name: '',
    email: '',
    message: '',
  });
  
  const submitForm = () => {
    console.log('Form submitted:', form.value);
  };
  </script>
  
  <style scoped>
  /* Center align labels and inputs */
  form .flex {
    align-items: center;
  }
  
  /* Button styles */
  .btn {
    padding: 12px;
    background: linear-gradient(to right, #ffcc00, #ff9900);
    color: white;
    font-weight: bold;
    border-radius: 25px;
    text-transform: uppercase;
    transition: transform 0.2s, box-shadow 0.2s;
    text-decoration: none;
    text-align: center;
  }
  
  .btn:hover {
    transform: scale(1.05);
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.3);
  }
  
  /* Ensure consistent vertical alignment */
  label {
    text-align: left;
    display: block;
    min-width: 9rem; /* Ensures enough space for labels */
  }
  
  input,
  textarea {
    flex: 1;
    padding: 12px;
    border: 1px solid #d1d5db; /* Neutral border color */
    border-radius: 8px;
    transition: border-color 0.3s, box-shadow 0.3s;
  }
  
  input:focus,
  textarea:focus {
    border-color: #fbbf24; /* Highlight border on focus */
    outline: none;
    box-shadow: 0 0 4px rgba(251, 191, 36, 0.5); /* Subtle glow effect */
  }
  </style>
  