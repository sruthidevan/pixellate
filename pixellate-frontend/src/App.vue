<template>
  <div id="app" class="flex flex-col min-h-screen">
    <!-- Navigation -->
    <nav class="bg-black text-white p-4">
      <ul class="flex justify-around">
        <li><router-link to="/" class="hover:text-yellow-300">Home</router-link></li>
        <li><router-link to="/services" class="hover:text-yellow-300">Services</router-link></li>
        <li><router-link to="/about" class="hover:text-yellow-300">About</router-link></li>
        <li><router-link to="/contact" class="hover:text-yellow-300">Contact</router-link></li>
      </ul>
    </nav>

    <!-- Main Content: Render the routed view here -->
    <div class="main-content flex-grow flex flex-col items-center justify-center">
      <router-view />
    </div>

    <!-- Footer -->
    <footer class="bg-black text-white text-center py-4 mt-8">
      <p>&copy; 2024 Pixellate. All rights reserved.</p>
    </footer>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
/* Add basic styles for navigation and layout */
nav ul {
  list-style: none;
  margin: 0;
  padding: 0;
  display: flex;
  justify-content: space-around;
}

nav ul li {
  display: inline;
}

nav ul li a {
  color: #ffffff;
  text-decoration: none;
  font-weight: bold;
  text-transform: uppercase;
}

nav ul li a:hover {
  color: #ffcc00;
}

.main-content {
  padding: 2rem;
  flex-grow: 1; /* Ensure main content takes up available space */
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

footer {
  border-top: 2px solid #ffcc00;
}
</style>