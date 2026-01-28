package generate

import (
	"bytes"
	"fmt"
)

var quranLists = []quranList{
	// arabic
	{id: "simple-clean", language: "Original", name: "Simple Clean", name_english: "Simple Clean", quran_type: "quran"},
	{id: "simple-min", language: "Original", name: "Simple Min", name_english: "Simple Min", quran_type: "quran"},
	{id: "simple-plain", language: "Original", name: "Simple Plain", name_english: "Simple Plain", quran_type: "quran"},
	{id: "simple", language: "Original", name: "Simple", name_english: "Simple", quran_type: "quran"},
	{id: "uthmani-min", language: "Original", name: "Uthmani Min", name_english: "Uthmani Min", quran_type: "quran"},
	{id: "uthmani", language: "Original", name: "Uthmani", name_english: "Uthmani", quran_type: "quran"},
	// translation
	{id: "am.sadiq", language: "Amharic", name: "ሳዲቅ & ሳኒ ሐቢብ", name_english: "Muhammed Sadiq and Muhammed Sani Habib", quran_type: "translation"},
	{id: "ar.jalalayn", language: "Arabic", name: "تفسير الجلالين", name_english: "Jalal ad-Din al-Mahalli and Jalal ad-Din as-Suyuti", quran_type: "translation"},
	{id: "ar.muyassar", language: "Arabic", name: "تفسير المیسر", name_english: "King Fahad Quran Complex", quran_type: "translation"},
	{id: "az.mammadaliyev", language: "Azerbaijani", name: "Məmmədəliyev & Bünyadov", name_english: "Vasim Mammadaliyev and Ziya Bunyadov", quran_type: "translation"},
	{id: "az.musayev", language: "Azerbaijani", name: "Musayev", name_english: "Alikhan Musayev", quran_type: "translation"},
	{id: "ber.mensur", language: "Amazigh", name: "At Mensur", name_english: "Ramdane At Mansour", quran_type: "translation"},
	{id: "bg.theophanov", language: "Bulgarian", name: "Теофанов", name_english: "Tzvetan Theophanov", quran_type: "translation"},
	{id: "bn.bengali", language: "Bengali", name: "মুহিউদ্দীন খান", name_english: "Muhiuddin Khan", quran_type: "translation"},
	{id: "bn.hoque", language: "Bengali", name: "জহুরুল হক", name_english: "Zohurul Hoque", quran_type: "translation"},
	{id: "bs.korkut", language: "Bosnian", name: "Korkut", name_english: "Besim Korkut", quran_type: "translation"},
	{id: "bs.mlivo", language: "Bosnian", name: "Mlivo", name_english: "Mustafa Mlivo", quran_type: "translation"},
	{id: "cs.hrbek", language: "Czech", name: "Hrbek", name_english: "Preklad I. Hrbek", quran_type: "translation"},
	{id: "cs.nykl", language: "Czech", name: "Nykl", name_english: "A. R. Nykl", quran_type: "translation"},
	{id: "de.aburida", language: "German", name: "Abu Rida", name_english: "Abu Rida Muhammad ibn Ahmad ibn Rassoul", quran_type: "translation"},
	{id: "de.bubenheim", language: "German", name: "Bubenheim & Elyas", name_english: "A. S. F. Bubenheim and N. Elyas", quran_type: "translation"},
	{id: "de.khoury", language: "German", name: "Khoury", name_english: "Adel Theodor Khoury", quran_type: "translation"},
	{id: "de.zaidan", language: "German", name: "Zaidan", name_english: "Amir Zaidan", quran_type: "translation"},
	{id: "dv.divehi", language: "Divehi", name: "ދިވެހި", name_english: "Office of the President of Maldives", quran_type: "translation"},
	{id: "en.ahmedali", language: "English", name: "Ahmed Ali", name_english: "Ahmed Ali", quran_type: "translation"},
	{id: "en.ahmedraza", language: "English", name: "Ahmed Raza Khan", name_english: "Ahmed Raza Khan", quran_type: "translation"},
	{id: "en.arberry", language: "English", name: "Arberry", name_english: "A. J. Arberry", quran_type: "translation"},
	{id: "en.daryabadi", language: "English", name: "Daryabadi", name_english: "Abdul Majid Daryabadi", quran_type: "translation"},
	{id: "en.hilali", language: "English", name: "Hilali & Khan", name_english: "Muhammad Taqi-ud-Din al-Hilali and Muhammad Muhsin Khan", quran_type: "translation"},
	{id: "en.itani", language: "English", name: "Itani", name_english: "Talal Itani", quran_type: "translation"},
	{id: "en.maududi", language: "English", name: "Maududi", name_english: "Abul Ala Maududi", quran_type: "translation"},
	{id: "en.mubarakpuri", language: "English", name: "Mubarakpuri", name_english: "Safi-ur-Rahman al-Mubarakpuri", quran_type: "translation"},
	{id: "en.pickthall", language: "English", name: "Pickthall", name_english: "Mohammed Marmaduke William Pickthall", quran_type: "translation"},
	{id: "en.qarai", language: "English", name: "Qarai", name_english: "Ali Quli Qarai", quran_type: "translation"},
	{id: "en.qaribullah", language: "English", name: "Qaribullah & Darwish", name_english: "Hasan al-Fatih Qaribullah and Ahmad Darwish", quran_type: "translation"},
	{id: "en.sahih", language: "English", name: "Saheeh International", name_english: "Saheeh International", quran_type: "translation"},
	{id: "en.sarwar", language: "English", name: "Sarwar", name_english: "Muhammad Sarwar", quran_type: "translation"},
	{id: "en.shakir", language: "English", name: "Shakir", name_english: "Mohammad Habib Shakir", quran_type: "translation"},
	{id: "en.transliteration", language: "English", name: "Transliteration", name_english: "English Transliteration", quran_type: "translation"},
	{id: "en.wahiduddin", language: "English", name: "Wahiduddin Khan", name_english: "Wahiduddin Khan", quran_type: "translation"},
	{id: "en.yusufali", language: "English", name: "Yusuf Ali", name_english: "Abdullah Yusuf Ali", quran_type: "translation"},
	{id: "es.bornez", language: "Spanish", name: "Bornez", name_english: "Raúl González Bórnez", quran_type: "translation"},
	{id: "es.cortes", language: "Spanish", name: "Cortes", name_english: "Julio Cortes", quran_type: "translation"},
	{id: "es.garcia", language: "Spanish", name: "Garcia", name_english: "Muhammad Isa García", quran_type: "translation"},
	{id: "ps.abdulwali", language: "Pashto", name: "عبدالولي", name_english: "Abdulwali Khan", quran_type: "translation"},
	{id: "fa.ansarian", language: "Persian", name: "انصاریان", name_english: "Hussain Ansarian", quran_type: "translation"},
	{id: "fa.ayati", language: "Persian", name: "آیتی", name_english: "AbdolMohammad Ayati", quran_type: "translation"},
	{id: "fa.bahrampour", language: "Persian", name: "بهرام‌پور", name_english: "Abolfazl Bahrampour", quran_type: "translation"},
	{id: "fa.fooladvand", language: "Persian", name: "فولادوند", name_english: "Mohammad Mahdi Fooladvand", quran_type: "translation"},
	{id: "fa.gharaati", language: "Persian", name: "قرائتی", name_english: "Mohsen Gharaati", quran_type: "translation"},
	{id: "fa.ghomshei", language: "Persian", name: "الهی قمشه‌ای", name_english: "Mahdi Elahi Ghomshei", quran_type: "translation"},
	{id: "fa.khorramdel", language: "Persian", name: "خرمدل", name_english: "Mostafa Khorramdel", quran_type: "translation"},
	{id: "fa.khorramshahi", language: "Persian", name: "خرمشاهی", name_english: "Bahaoddin Khorramshahi", quran_type: "translation"},
	{id: "fa.makarem", language: "Persian", name: "مکارم شیرازی", name_english: "Naser Makarem Shirazi", quran_type: "translation"},
	{id: "fa.moezzi", language: "Persian", name: "معزی", name_english: "Mohammad Kazem Moezzi", quran_type: "translation"},
	{id: "fa.mojtabavi", language: "Persian", name: "مجتبوی", name_english: "Sayyed Jalaloddin Mojtabavi", quran_type: "translation"},
	{id: "fa.sadeqi", language: "Persian", name: "صادقی تهرانی", name_english: "Mohammad Sadeqi Tehrani", quran_type: "translation"},
	{id: "fa.safavi", language: "Persian", name: "صفوی", name_english: "Sayyed Mohammad Reza Safavi", quran_type: "translation"},
	{id: "fr.hamidullah", language: "French", name: "Hamidullah", name_english: "Muhammad Hamidullah", quran_type: "translation"},
	{id: "ha.gumi", language: "Hausa", name: "Gumi", name_english: "Abubakar Mahmoud Gumi", quran_type: "translation"},
	{id: "hi.farooq", language: "Hindi", name: "फ़ारूक़ ख़ान & अहमद", name_english: "Muhammad Farooq Khan and Muhammad Ahmed", quran_type: "translation"},
	{id: "hi.hindi", language: "Hindi", name: "फ़ारूक़ ख़ान & नदवी", name_english: "Suhel Farooq Khan and Saifur Rahman Nadwi", quran_type: "translation"},
	{id: "id.indonesian", language: "Indonesian", name: "Bahasa Indonesia", name_english: "Indonesian Ministry of Religious Affairs", quran_type: "translation"},
	{id: "id.jalalayn", language: "Indonesian", name: "Tafsir Jalalayn", name_english: "Jalal ad-Din al-Mahalli and Jalal ad-Din as-Suyuti", quran_type: "translation"},
	{id: "id.muntakhab", language: "Indonesian", name: "Quraish Shihab", name_english: "Muhammad Quraish Shihab et al.", quran_type: "translation"},
	{id: "it.piccardo", language: "Italian", name: "Piccardo", name_english: "Hamza Roberto Piccardo", quran_type: "translation"},
	{id: "ja.japanese", language: "Japanese", name: "Japanese", name_english: "Unknown", quran_type: "translation"},
	{id: "ko.korean", language: "Korean", name: "Korean", name_english: "Unknown", quran_type: "translation"},
	{id: "ku.asan", language: "Kurdish", name: "ته‌فسیری ئاسان", name_english: "Burhan Muhammad-Amin", quran_type: "translation"},
	{id: "ml.abdulhameed", language: "Malayalam", name: "അബ്ദുല്‍ ഹമീദ് & പറപ്പൂര്‍", name_english: "Cheriyamundam Abdul Hameed and Kunhi Mohammed Parappoor", quran_type: "translation"},
	{id: "ml.karakunnu", language: "Malayalam", name: "കാരകുന്ന് & എളയാവൂര്", name_english: "Muhammad Karakunnu and Vanidas Elayavoor", quran_type: "translation"},
	{id: "ms.basmeih", language: "Malay", name: "Basmeih", name_english: "Abdullah Muhammad Basmeih", quran_type: "translation"},
	{id: "nl.keyzer", language: "Dutch", name: "Keyzer", name_english: "Salomo Keyzer", quran_type: "translation"},
	{id: "nl.leemhuis", language: "Dutch", name: "Leemhuis", name_english: "Fred Leemhuis", quran_type: "translation"},
	{id: "nl.siregar", language: "Dutch", name: "Siregar", name_english: "Sofian S. Siregar", quran_type: "translation"},
	{id: "no.berg", language: "Norwegian", name: "Einar Berg", name_english: "Einar Berg", quran_type: "translation"},
	{id: "pl.bielawskiego", language: "Polish", name: "Bielawskiego", name_english: "Józefa Bielawskiego", quran_type: "translation"},
	{id: "pt.elhayek", language: "Portuguese", name: "El-Hayek", name_english: "Samir El-Hayek", quran_type: "translation"},
	{id: "ro.grigore", language: "Romanian", name: "Grigore", name_english: "George Grigore", quran_type: "translation"},
	{id: "ru.abuadel", language: "Russian", name: "Абу Адель", name_english: "Abu Adel", quran_type: "translation"},
	{id: "ru.kalam", language: "Russian", name: "Калям Шариф", name_english: "Muslim Religious Board of the Republiс of Tatarstan", quran_type: "translation"},
	{id: "ru.krachkovsky", language: "Russian", name: "Крачковский", name_english: "Ignaty Yulianovich Krachkovsky", quran_type: "translation"},
	{id: "ru.kuliev-alsaadi", language: "Russian", name: "Кулиев + ас-Саади", name_english: "Elmir Kuliev (with Abd ar-Rahman as-Saadi&apos;s commentaries)", quran_type: "translation"},
	{id: "ru.kuliev", language: "Russian", name: "Кулиев", name_english: "Elmir Kuliev", quran_type: "translation"},
	{id: "ru.muntahab", language: "Russian", name: "Аль-Мунтахаб", name_english: "Ministry of Awqaf, Egypt", quran_type: "translation"},
	{id: "ru.osmanov", language: "Russian", name: "Османов", name_english: "Magomed-Nuri Osmanovich Osmanov", quran_type: "translation"},
	{id: "ru.porokhova", language: "Russian", name: "Порохова", name_english: "V. Porokhova", quran_type: "translation"},
	{id: "ru.sablukov", language: "Russian", name: "Саблуков", name_english: "Gordy Semyonovich Sablukov", quran_type: "translation"},
	{id: "sd.amroti", language: "Sindhi", name: "امروٽي", name_english: "Taj Mehmood Amroti", quran_type: "translation"},
	{id: "so.abduh", language: "Somali", name: "Abduh", name_english: "Mahmud Muhammad Abduh", quran_type: "translation"},
	{id: "sq.ahmeti", language: "Albanian", name: "Sherif Ahmeti", name_english: "Sherif Ahmeti", quran_type: "translation"},
	// {id: "sq.mehdiu", language: "Albanian", name: "Feti Mehdiu", name_english: "Feti Mehdiu", quran_type: "translation"},
	{id: "sq.nahi", language: "Albanian", name: "Efendi Nahi", name_english: "Hasan Efendi Nahi", quran_type: "translation"},
	{id: "sv.bernstrom", language: "Swedish", name: "Bernström", name_english: "Knut Bernström", quran_type: "translation"},
	{id: "sw.barwani", language: "Swahili", name: "Al-Barwani", name_english: "Ali Muhsin Al-Barwani", quran_type: "translation"},
	{id: "ta.tamil", language: "Tamil", name: "ஜான் டிரஸ்ட்", name_english: "Jan Turst Foundation", quran_type: "translation"},
	{id: "tg.ayati", language: "Tajik", name: "Оятӣ", name_english: "AbdolMohammad Ayati", quran_type: "translation"},
	{id: "th.thai", language: "Thai", name: "ภาษาไทย", name_english: "King Fahad Quran Complex", quran_type: "translation"},
	{id: "tr.ates", language: "Turkish", name: "Süleyman Ateş", name_english: "Suleyman Ates", quran_type: "translation"},
	{id: "tr.bulac", language: "Turkish", name: "Alİ Bulaç", name_english: "Alİ Bulaç", quran_type: "translation"},
	{id: "tr.diyanet", language: "Turkish", name: "Diyanet İşleri", name_english: "Diyanet Isleri", quran_type: "translation"},
	{id: "tr.golpinarli", language: "Turkish", name: "Abdulbakî Gölpınarlı", name_english: "Abdulbaki Golpinarli", quran_type: "translation"},
	{id: "tr.ozturk", language: "Turkish", name: "Öztürk", name_english: "Yasar Nuri Ozturk", quran_type: "translation"},
	{id: "tr.transliteration", language: "Turkish", name: "Çeviriyazı", name_english: "Muhammet Abay", quran_type: "translation"},
	{id: "tr.vakfi", language: "Turkish", name: "Diyanet Vakfı", name_english: "Diyanet Vakfi", quran_type: "translation"},
	{id: "tr.yazir", language: "Turkish", name: "Elmalılı Hamdi Yazır", name_english: "Elmalili Hamdi Yazir", quran_type: "translation"},
	{id: "tr.yildirim", language: "Turkish", name: "Suat Yıldırım", name_english: "Suat Yildirim", quran_type: "translation"},
	{id: "tr.yuksel", language: "Turkish", name: "Edip Yüksel", name_english: "Edip Yüksel", quran_type: "translation"},
	{id: "tt.nugman", language: "Tatar", name: "Yakub Ibn Nugman", name_english: "Yakub Ibn Nugman", quran_type: "translation"},
	{id: "ug.saleh", language: "Uyghur", name: "محمد صالح", name_english: "Muhammad Saleh", quran_type: "translation"},
	{id: "ur.ahmedali", language: "Urdu", name: "احمد علی", name_english: "Ahmed Ali", quran_type: "translation"},
	{id: "ur.jalandhry", language: "Urdu", name: "جالندہری", name_english: "Fateh Muhammad Jalandhry", quran_type: "translation"},
	{id: "ur.jawadi", language: "Urdu", name: "علامہ جوادی", name_english: "Syed Zeeshan Haider Jawadi", quran_type: "translation"},
	{id: "ur.junagarhi", language: "Urdu", name: "محمد جوناگڑھی", name_english: "Muhammad Junagarhi", quran_type: "translation"},
	{id: "ur.kanzuliman", language: "Urdu", name: "احمد رضا خان", name_english: "Ahmed Raza Khan", quran_type: "translation"},
	{id: "ur.maududi", language: "Urdu", name: "ابوالاعلی مودودی", name_english: "Abul Ala Maududi", quran_type: "translation"},
	{id: "ur.najafi", language: "Urdu", name: "محمد حسین نجفی", name_english: "Muhammad Hussain Najafi", quran_type: "translation"},
	{id: "ur.qadri", language: "Urdu", name: "طاہر القادری", name_english: "Tahir ul Qadri", quran_type: "translation"},
	{id: "uz.sodik", language: "Uzbek", name: "Мухаммад Содик", name_english: "Muhammad Sodik Muhammad Yusuf", quran_type: "translation"},
	{id: "zh.jian", language: "Chinese", name: "Ma Jian", name_english: "Ma Jian", quran_type: "translation"},
	{id: "zh.majian", language: "Chinese", name: "Ma Jian (Traditional)", name_english: "Ma Jian", quran_type: "translation"},
}

type quranList struct {
	id           string
	language     string
	name         string
	name_english string
	quran_type   string
}

func Quran() {
	var b bytes.Buffer
	b.WriteString(`/************************************************** quran **************************************************/
insert ignore into quran(id, language, name, name_english, quran_type) values
`)
	for i := 0; i < len(quranLists)-1; i++ {
		var item = quranLists[i]
		b.WriteString(fmt.Sprintf("('%s', '%s', '%s', '%s', '%s'),\n", item.id, item.language, item.name, item.name_english, item.quran_type))
	}
	// add the last quran with a semi-colon instead of comma and new line
	var item = quranLists[len(quranLists)-1]
	b.WriteString(fmt.Sprintf("('%s', '%s', '%s', '%s', '%s');\n", item.id, item.language, item.name, item.name_english, item.quran_type))

	// write to file
	appendFile(dataload, b.String())
}
