#!/bin/bash

################################################## variables ##################################################
qurans_basic=(
  simple
)
qurans_with_pause_marks=(
  simple-clean
  simple-min
  simple-plain
  uthmani
  uthmani-min
)
translations=(
  am.sadiq
  ar.jalalayn
  ar.muyassar
  az.mammadaliyev
  az.musayev
  ber.mensur
  bg.theophanov
  bn.bengali
  bn.hoque
  bs.korkut
  bs.mlivo
  cs.hrbek
  cs.nykl
  de.aburida
  de.bubenheim
  de.khoury
  de.zaidan
  dv.divehi
  en.ahmedali
  en.ahmedraza
  en.arberry
  en.daryabadi
  en.hilali
  en.itani
  en.maududi
  en.mubarakpuri
  en.pickthall
  en.qarai
  en.qaribullah
  en.sahih
  en.sarwar
  en.shakir
  en.transliteration
  en.wahiduddin
  en.yusufali
  es.bornez
  es.cortes
  es.garcia
  fa.ansarian
  fa.ayati
  fa.bahrampour
  fa.fooladvand
  fa.gharaati
  fa.ghomshei
  fa.khorramdel
  fa.khorramshahi
  fa.makarem
  fa.moezzi
  fa.mojtabavi
  fa.sadeqi
  fa.safavi
  fr.hamidullah
  ha.gumi
  hi.farooq
  hi.hindi
  id.indonesian
  id.jalalayn
  id.muntakhab
  it.piccardo
  ja.japanese
  ko.korean
  ku.asan
  ml.abdulhameed
  ml.karakunnu
  ms.basmeih
  nl.keyzer
  nl.leemhuis
  nl.siregar
  no.berg
  pl.bielawskiego
  ps.abdulwali
  pt.elhayek
  ro.grigore
  ru.abuadel
  ru.kalam
  ru.krachkovsky
  ru.kuliev
  ru.kuliev-alsaadi
  ru.muntahab
  ru.osmanov
  ru.porokhova
  ru.sablukov
  sd.amroti
  so.abduh
  sq.ahmeti
  sq.mehdiu
  sq.nahi
  sv.bernstrom
  sw.barwani
  ta.tamil
  tg.ayati
  th.thai
  tr.ates
  tr.bulac
  tr.diyanet
  tr.golpinarli
  tr.ozturk
  tr.transliteration
  tr.vakfi
  tr.yazir
  tr.yildirim
  tr.yuksel
  tt.nugman
  ug.saleh
  ur.ahmedali
  ur.jalandhry
  ur.jawadi
  ur.junagarhi
  ur.kanzuliman
  ur.maududi
  ur.najafi
  ur.qadri
  uz.sodik
  zh.jian
  zh.majian
)

################################################## methods ##################################################
download_qurans_basic() {
  local list=("$@") # get an array for rest of the arguments
  for item in ${list[@]}; do
    echo $item
    curl --silent "https://tanzil.net/pub/download/index.php?quranType=$item&sajdah=true&outType=txt" --output "tanzil/$item.txt"
  done
}

download_qurans_with_pause_marks() {
  local list=("$@") # get an array for rest of the arguments
  for item in ${list[@]}; do
    echo $item
    curl --silent "https://tanzil.net/pub/download/index.php?quranType=$item&marks=true&sajdah=true&outType=txt" --output "tanzil/$item.txt"
  done
}

download_translations() {
  local translationList=("$@") # get an array for rest of the arguments
  for translation in ${translationList[@]}; do
    echo $translation
    curl --silent "https://tanzil.net/trans/?transID=$translation&type=txt" --output "tanzil/$translation.txt"
  done
}

################################################## execution ##################################################

mkdir tanzil -p
curl --silent -H 'accept:application/xml;' "https://tanzil.net/res/text/metadata/quran-data.xml" --output "tanzil/quran-data.xml"
download_qurans_basic ${qurans_basic[@]}
download_qurans_with_pause_marks ${qurans_with_pause_marks[@]}
download_translations ${translations[@]}
