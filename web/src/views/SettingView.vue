<script setup>
import { ref, onMounted } from 'vue'
import { inject } from 'vue'
import { QuranStore } from "@/store/index.js";

const apiClient = inject('apiClient');
const dbClient = inject('dbClient');
const storageClient = inject('storageClient');

const setting = ref({});
const qurans = ref([]);

onMounted(async () => {
  setting.value = storageClient.getSetting();
  await dbClient.open()
  qurans.value = await new QuranStore(dbClient, apiClient).getAll()
});

function getName(quranId) {
  let result = qurans.value.find(quran => quran.id == quranId)
  if (result == null) {
    return "";
  }
  else if (result.quran_type == "quran") {
    return "بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ"
  }
  else {
    return result == null ? "" : result.name;
  }
}

function textStyle(index) {
  let size = index == 0 ? setting.value.selectedQuranSize1 : setting.value.selectedQuranSize2;
  return `font-size: ${size}px;`
}

const save = () => {
  storageClient.setSetting(setting.value);
}
const saveAndClearStore = () => {
  storageClient.setSetting(setting.value);
  dbClient.clearStore('quranLine');
}

</script>

<template>
  <div class="setting">

    <select v-model="setting.selectedQuran1" @change="saveAndClearStore">
      <option v-for="quran in qurans" :key="quran.id" :value="quran.id">{{ quran.language }} : {{ quran.name }}</option>
    </select>
    <input v-model="setting.selectedQuranSize1" type="range" min="2" max="100" step="2" style="width:100%;"
      @change="save">
    <span :style="textStyle(0)">{{ getName(setting.selectedQuran1) }}</span>
    <select v-model="setting.selectedQuran2" @change="saveAndClearStore">
      <option v-for="quran in qurans" :key="quran.id" :value="quran.id">{{ quran.language }} : {{ quran.name }}</option>
    </select>
    <span :style="textStyle(1)">{{ getName(setting.selectedQuran2) }}</span>
    <input v-model="setting.selectedQuranSize2" type="range" min="2" max="100" step="2" style="width:100%;"
      @change="save">
    <select v-model="setting.paging" @change="save">
      <option value="page">Page</option>
      <option value="surah">Surah</option>
      <option value="ruku">Ruku</option>
      <option value="hizb">Hizb</option>
      <option value="juz">Juz</option>
      <option value="manzil">Manzil</option>
    </select>
    <router-link to="/quran"><button role="link">Save</button></router-link>
  </div>

</template>

<style scoped>
.setting {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100%;
  padding: 20px;
  gap: 10px;
}
</style>
