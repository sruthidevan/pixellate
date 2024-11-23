<template>
    <div class="min-h-screen flex items-center justify-center bg-gradient-to-b from-yellow-100 via-yellow-200 to-yellow-300">
      <div class="bg-white p-8 rounded-lg shadow-md w-full max-w-lg text-center">
        <h1 class="text-3xl font-bold mb-6 text-black">Contact Us</h1>
        <p class="text-lg mb-6 text-gray-700">
          You can reach us at:
        </p>
        <p class="text-lg font-bold text-gray-900">
          <a href="mailto:antilles1990@gmail.com" class="text-blue-500 hover:underline">antilles1990@gmail.com</a>
        </p>
      </div>
    </div>
  </template>
  
  <script setup>
  // No custom logic needed for this static page.
  </script>
  
  <style scoped>
  /* Center align the content */
  div {
    text-align: center;
  }
  
  /* Styling for the email text */
  p {
    margin: 0;
    padding: 0;
  }
  
  /* Styling for the mailto link */
  a {
    color: #1d4ed8; /* Blue color */
    text-decoration: none;
  }
  
  a:hover {
    text-decoration: underline;
  }
  </style>