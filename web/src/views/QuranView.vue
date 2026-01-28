<script setup>
import { ref, onMounted, computed } from 'vue'
import { inject } from 'vue'
import { SuraStore, QuranLineStore } from '@/store';

const apiClient = inject('apiClient');
const dbClient = inject('dbClient');
const storageClient = inject('storageClient');

const setting = ref({});
const currentPage = ref({});
const quranLines = ref([]);
const suras = ref([]);
const pageMax = ref(604);

onMounted(async () => {
  setting.value = storageClient.getSetting();
  currentPage.value = storageClient.getCurrentPage();
  await dbClient.open()
  quranLines.value = await new QuranLineStore(dbClient, apiClient).getByKey(setting.value.selectedQuran1, setting.value.selectedQuran2, setting.value.paging, currentPage.value)
  suras.value = await new SuraStore(dbClient, apiClient).getAll();
});

async function displayPage() {
  storageClient.setCurrentPage(currentPage.value);
  quranLines.value = await new QuranLineStore(dbClient, apiClient).getByKey(setting.value.selectedQuran1, setting.value.selectedQuran2, setting.value.paging, currentPage.value)
  window.scrollTo(0, 0)
}

function currentSura(quranLine) {
  let result = suras.value.find(sura => sura.id == quranLine.surah && quranLine.aya == 1)
  return result == null ? null : `${result.name_arabic} ${result.name_english} (${result.id})`;
};

function showBismillah(quranLine) {
  let result = suras.value.find(sura => sura.id == quranLine.surah && quranLine.aya == 1)
  return (result != null && result.id != 1 && result.id != 9)
}

function textStyle(index) {
  var style = "";
  var selectedQuran = setting.value[`selectedQuran${index + 1}`];
  let size = index == 0 ? setting.value.selectedQuranSize1 : setting.value.selectedQuranSize2;
  switch (selectedQuran) {
    case 'simple':
      style += `line-height: ${Number(size) + 10}px;`
      break;
    default:
      style += `line-height: ${Number(size) + 0}px;`
      break
  }
  return `${style}font-size: ${size}px;${style}`
}

const sortedItems = computed(() => {
  // Sort by line
  const sorted = quranLines.value.sort((a, b) => a.line - b.line);

  const grouped = sorted.reduce((acc, item) => {
    if (!acc[item.line]) {
      acc[item.line] = []
    }
    acc[item.line].push(item);
    return acc;
  }, {});
  return grouped;
});



</script>
<template>
  <div id="quran">
    <header>
      <div></div>
      <div>
        <input type="button" value="<" @click="currentPage++; displayPage()"></input>
        <input v-model="currentPage" type="number" min="1" :max="pageMax" @change="displayPage" class="pageNumber" />
        <input type="button" value=">" @click="currentPage--; displayPage()"></input>
      </div>
      <RouterLink to="/setting">
        <svg id="btnSetting" viewBox="0 0 512 512" width="24" height="24" fill="silver">
          <path
            d="M487.4 315.7l-42.6-24.6c4.3-23.2 4.3-47 0-70.2l42.6-24.6c4.9-2.8 7.1-8.6 5.5-14-11.1-35.6-30-67.8-54.7-94.6-3.8-4.1-10-5.1-14.8-2.3L380.8 110c-17.9-15.4-38.5-27.3-60.8-35.1V25.8c0-5.6-3.9-10.5-9.4-11.7-36.7-8.2-74.3-7.8-109.2 0-5.5 1.2-9.4 6.1-9.4 11.7V75c-22.2 7.9-42.8 19.8-60.8 35.1L88.7 85.5c-4.9-2.8-11-1.9-14.8 2.3-24.7 26.7-43.6 58.9-54.7 94.6-1.7 5.4.6 11.2 5.5 14L67.3 221c-4.3 23.2-4.3 47 0 70.2l-42.6 24.6c-4.9 2.8-7.1 8.6-5.5 14 11.1 35.6 30 67.8 54.7 94.6 3.8 4.1 10 5.1 14.8 2.3l42.6-24.6c17.9 15.4 38.5 27.3 60.8 35.1v49.2c0 5.6 3.9 10.5 9.4 11.7 36.7 8.2 74.3 7.8 109.2 0 5.5-1.2 9.4-6.1 9.4-11.7v-49.2c22.2-7.9 42.8-19.8 60.8-35.1l42.6 24.6c4.9 2.8 11 1.9 14.8-2.3 24.7-26.7 43.6-58.9 54.7-94.6 1.5-5.5-.7-11.3-5.6-14.1zM256 336c-44.1 0-80-35.9-80-80s35.9-80 80-80 80 35.9 80 80-35.9 80-80 80z">
          </path>
        </svg>
      </RouterLink>
    </header>
    <main>
      <div v-for="(items, line) in sortedItems" :key="line">
        <div v-if="currentSura(items[0])" class="ayah" :style="textStyle(0)">
          <div>{{ currentSura(items[0]) }}</div>
          <div style="padding-bottom: 8px" v-if="showBismillah(items[0])">بِسْمِ اللَّهِ الرَّحْمَٰنِ الرَّحِيمِ</div>
        </div>
        <div class="ayah">
          <div v-for="(item, i) in items" :key="item.quran_id">
            <span :style="textStyle(i)" v-html="item.text"></span>
            <span v-if="i == items.length - 1" class="pageNumber"> ({{ items[0].aya }}) </span>
          </div>
        </div>
      </div>
    </main>
    <footer>
      <input type="button" value="<" @click="currentPage++; displayPage()"></input>
      <input v-model="currentPage" type="number" min="1" :max="pageMax" @change="displayPage" class="pageNumber" />
      <input type="button" value=">" @click="currentPage--; displayPage()"></input>
    </footer>
  </div>
</template>

<style scoped>
header,
footer {
  text-align: center;
  margin: 0px 0px;
}

header {
  display: flex;
  flex-direction: row;
  justify-content: center;
  justify-content: space-between;

}

.ayah {
  text-align: center;
  background-color: black;
  border-radius: 10px;
  margin: 0px 0px 3px 0px;
  padding: 0px 3px;
}

input {
  width: 50px;
  height: 25px;
  padding: 0px;
  margin: 0px;
  color: #ffffff;
  background-color: #000000;
  align-self: center;
  text-align: center;
  /* font-size: x-small; */
}
</style>
