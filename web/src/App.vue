<script setup>
import { onMounted, inject, ref } from 'vue';
import { RouterLink, RouterView } from 'vue-router'

const apiClient = inject('apiClient');
const dbClient = inject('dbClient');

onMounted(async () => {
  await dbClient.open();
});

</script>

<template>
  <RouterView />
</template>

<style scoped></style>
