package crawl

import (
	"encoding/json"
	"fmt"
	"mxgk/crawl/models"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/net/html"
)

type VideoItem struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func CrawlVideo() {
	// ///lớp 6
	// crawlVideo("PLCd8j6ZYo0la73mUTNbNjCx2Unpkgfuyw", "g6")
	// crawlVideo("PLCd8j6ZYo0lbZqAdLOSkjXgKGGz_2Taf5", "g6")
	// crawlVideo("PLCd8j6ZYo0lbyefvfLKgpoGU9DnSOsvpz", "g6")
	// crawlVideo("PLCd8j6ZYo0lb1MNlwtvKn8Po6Nw-DQDim", "g6")
	// crawlVideo("PLCd8j6ZYo0lYJ5jqIhAZV5NWqftQF41Ag", "g6")
	// crawlVideo("PL0g6dPJJizVbUjIL_CjNZHCzznaG9iYV3", "g6")
	// crawlVideo("PL0g6dPJJizVaFjeFQJaZX4GwxJj-xPjKU", "g6")
	// crawlVideo("PL0g6dPJJizVasQSqDR6GAiMUesxu1j082", "g6")
	// crawlVideo("PL0g6dPJJizVYSq5Q0p_osdvKgD3tdV9Sv", "g6")
	// crawlVideo("PLPUZuj69QZud9L_mrUwV5yEAMDqhbract", "g6")
	// crawlVideo("PLzS2cTFGpDyxOvV83mmcGltpA_MmUTt_m", "g6")
	// crawlVideo("PLzS2cTFGpDyxtJijBo4MbkBK_dsewdLeN", "g6")
	// crawlVideo("PLzS2cTFGpDyzSlfwHMwspdmC6qyY9ULCE", "g6")
	// crawlVideo("PLzS2cTFGpDyxRKQUH-N_e32HUM0OzLnkx", "g6")
	// crawlVideo("PLzS2cTFGpDyyJle9dUJPD8Vem9fzDiV6t", "g6")
	// crawlVideo("PLzS2cTFGpDyzNjUy5f5WQ-x8-8yZP2M29", "g6")
	// crawlVideo("PLzS2cTFGpDyxFSmMmlgSru4UY_DLgmwI-", "g6")
	// crawlVideo("PLzS2cTFGpDywqopA_dZScuefq5eSFJ3cK", "g6")
	// crawlVideo("PLzS2cTFGpDywC4v_kPqrPkRT1_WreFzS5", "g6")
	// crawlVideo("PLzS2cTFGpDyyGPfJ9f8TNA7oIxuZzbi-4", "g6")
	// crawlVideo("PLzS2cTFGpDyySOK9bqm9AW5bJc1fmDQ3P", "g6")
	// crawlVideo("PLzS2cTFGpDyy4dJvSbyYsaorDuA0TbDZN", "g6")
	// crawlVideo("PLzS2cTFGpDyzh-0ZSN8ljpC1zB7n5Ktqg", "g6")
	// crawlVideo("PLxUdsmN71otAxy40bm5mMHJRUC9p_GF9C", "g6")
	// crawlVideo("PL2Af1ctPF5GMtPsUMtmBvJYKAGQA9JaQW", "g6")
	// crawlVideo("PL2Af1ctPF5GNffgRz3Mk6cVppqZwYE0Wp", "g6")
	// crawlVideo("PLFV0Zv0oRuG6WvDzD69hhoJPSnNH2HWhn", "g6")
	// crawlVideo("PLFV0Zv0oRuG5T36K7MjlN1oYzucciwOWh", "g6")
	// crawlVideo("PLFV0Zv0oRuG6SQyu2QFx9rfypeJVWt_dk", "g6")
	// crawlVideo("PLFV0Zv0oRuG76xe5RIa8gi8I-vKZiBIWb", "g6")
	// crawlVideo("PLFV0Zv0oRuG6QSbUgxPuFry7cY3ZQtRYP", "g6")
	// crawlVideo("PLFV0Zv0oRuG4rNdTsTwfgTf5mNteBQ1Bj", "g6")
	// crawlVideo("PLFV0Zv0oRuG6BjKlVmZVFeRV0xcW9XzR0", "g6")
	// crawlVideo("PLHlw36ksrsNlg38fQKvIdyqPa3caGvCnA", "g6")
	// crawlVideo("PLHlw36ksrsNlPpSb--b88bmzwszp1K5df", "g6")
	// ///lớp 7
	// crawlVideo("PLHlw36ksrsNljzwxN6Q6rJzSDkOFLnK7L", "g7")
	// crawlVideo("PLqcCOZrA1NA2svYLMS4riatEG04G-Vggk", "g7")
	// crawlVideo("PLxUdsmN71otBzaOoOBKHdvuBEghSQnm3x", "g7")
	// crawlVideo("PLCd8j6ZYo0lY2vB8_kxXBQ2rj0_6b0RVD", "g7")
	// crawlVideo("PLCd8j6ZYo0lbFhIEuEF8yU4IKQJUK9OaY", "g7")
	// crawlVideo("PLxUdsmN71otAhZrSjHhFdgE08GXJ3VLL3", "g7")
	// crawlVideo("PLqcCOZrA1NA2ai4HkgkUovNG_IV07K9wb", "g7")
	// crawlVideo("PLMzRq608THSM5moWUfAXtLyezvj9gBmO2", "g7")
	// crawlVideo("PLHlw36ksrsNnzNxDTw5tp-QhYBWBUGF0v", "g7")
	// crawlVideo("PL-0HTRU3ATXBgsSemNtNcq8MqqbrEA27D", "g7")
	// crawlVideo("PLDHr_ecbSve6OQagl6IK3AafNuXNi2UEb", "g7")
	// crawlVideo("PLCd8j6ZYo0lae6v6H3Z0HpCRave-IDmMi", "g7")
	// crawlVideo("PLqcCOZrA1NA2z7wTvdKxct0MXciHyZx4k", "g7")
	// crawlVideo("PL7lt8bNTm9pzVSjoQIiHEPlW6jr20d8hK", "g7")
	// crawlVideo("PLzdVwemHYxxu7fQuPEaIU-7BvWE3XJSv2", "g7")
	// crawlVideo("PLYmriOlxq-miYZI5IKjT59f2sJLGzdvJl", "g7")
	// crawlVideo("PLCd8j6ZYo0lbZpOTH-AvpljZqe3rAnvQy", "g7")
	// crawlVideo("PLqcCOZrA1NA2ZhdtckMtqr351UExmeT_G", "g7")
	// crawlVideo("PLA33KQsrdG3jyLsBKy3b-FB3qMBRAjryW", "g7")
	// crawlVideo("PL3yjsAgvkjdTTpwLQYhWzOlLyjtpAm1zh", "g7")
	// crawlVideo("PLxUdsmN71otC0yo-yTDk4Q6685ZRcA0L3", "g7")
	// crawlVideo("PLqcCOZrA1NA3nFRiFStJWZThgb3IArkZO", "g7")
	// crawlVideo("PLxUdsmN71otDcV3TJ4LTWSlZNwcSMbowW", "g7")
	// crawlVideo("PLzS2cTFGpDyypMEu5nPIjr0TYTWdw_fnZ", "g7")
	// crawlVideo("PLCd8j6ZYo0lYp2u8igDarK_gOq3AZZ_xI", "g7")
	// crawlVideo("PLzS2cTFGpDyxb9q5iK4KRdPjuz1nGauaq", "g7")
	// crawlVideo("PL-0HTRU3ATXDQNQERORdc23fqzPcX6Sh2", "g7")
	// crawlVideo("PLxUdsmN71otDoOO_Lfp1aNN7B93RiEIL8", "g7")
	// crawlVideo("PLb86fQcyLH4QXyd__t53ISIAJbaWaT73Z", "g7")
	// crawlVideo("PLzS2cTFGpDyyPm-s6XzsMUpDDrfBnshjh", "g7")
	// crawlVideo("PL3yjsAgvkjdTzxX3JBRr7PiJArM0GlUE6", "g7")
	// crawlVideo("PLDHr_ecbSve4Nqe9UL5vJXpnZWQI9Yp5J", "g7")
	// crawlVideo("PLzS2cTFGpDyw0RJIMbrvM6x6gixHNqTa7", "g7")
	// crawlVideo("PLDHr_ecbSve7SZ-A2FF4jSAjJKGoURur3", "g7")
	// crawlVideo("PLqcCOZrA1NA2CLpRIgnGxKLdV1sifLAqR", "g7")
	// crawlVideo("PL5kxOdJ5P408G2IIaHy3sDdQn8RPo3ek6", "g7")
	// crawlVideo("PLzS2cTFGpDyz4x5d0zNvX0-LcA83dRPnU", "g7")
	// crawlVideo("PLqcCOZrA1NA02dN22PMnHLZeQ_MHj2ZhP", "g7")
	// crawlVideo("PL5kxOdJ5P409OAjjgKUNqVc7zP_1YHNPQ", "g7")
	// crawlVideo("PL-qYbpxhuvCOm9ty__q3qY7eXAfgJNFlf", "g7")
	// ///lớp 8
	// crawlVideo("PLzS2cTFGpDywrRoynLkUo0uux_KTKnsAK", "g8")
	// crawlVideo("PLhUgn84MlNfbE9YeNNzzk2xez2zDR1p70", "g8")
	// crawlVideo("PL8P0lAsvM4EufMXyGHhr7-VFbvacOjvd_", "g8")
	// crawlVideo("PLzS2cTFGpDyzi7o5Aj0fW7QLPPrBVpP91", "g8")
	// crawlVideo("PL5kxOdJ5P40_5vRTyqABP40WjwaZJ1Vb-", "g8")
	// crawlVideo("PL3jYV5GzVcJB20ZYlXDbg4DYNiDTM11nk", "g8")
	// crawlVideo("PL7udFEgGtyYW1Ep83Vaon86KWQtEDMxJ-", "g8")
	// crawlVideo("PLXPgM84qwrZAZzsj_jz-hIG6xxVaEz-XQ", "g8")
	// crawlVideo("PLuMGE4ip9nTVi25Vs1N3lTHK0BT1liYhk", "g8")
	// crawlVideo("PLhUgn84MlNfbPkmPasJamfd0R9r24s1sN", "g8")
	// crawlVideo("PLv_FkS88YBCfYI1HWUeK9KJ9YH24nF64C", "g8")
	// crawlVideo("PL5kxOdJ5P40_oTs5V5HLk8nPE_YTHw6oq", "g8")
	// crawlVideo("PL6_AM0jq2YzqF297k9boCxAwN6YprzkFF", "g8")
	// crawlVideo("PLLl22mzbOTrWj_-rRJBTTK3Y2JZyqHSRE", "g8")
	// crawlVideo("PLs_rMAIUyBHjMPHSo76-pJNO-Lq_BMgYZ", "g8")
	// crawlVideo("PLMzRq608THSOwIGP9N4gOvyz9I_bXFFwG", "g8")
	// crawlVideo("PLuMGE4ip9nTVRr3VHdTxZKDHT44oB9PGe", "g8")
	// crawlVideo("PLuMGE4ip9nTUpNpFM5fL13grKV56VYcpW", "g8")
	// crawlVideo("PLF02NpP4n1QAwlMTLOzDs8RckK6N-whUB", "g8")
	// crawlVideo("PLzS2cTFGpDyyyNE0UDrcJQT8WEhhznqIX", "g8")
	// crawlVideo("PL5kxOdJ5P40_u1BL9RN5-PMctii0G0DrH", "g8")
	// crawlVideo("PL5kxOdJ5P408w10ahflbWwIBqmtLttCpx", "g8")
	// crawlVideo("PLidK1VaE4fKdGgJcDQP6w5aWKDSSfphVm", "g8")
	// crawlVideo("PLQimOW2Ay0xbHPDNpVeXqLmvQEOPb6vOb", "g8")
	// crawlVideo("PLqcCOZrA1NA3CkmY5ngRra_rKaIvBg1I0", "g8")
	// crawlVideo("PL7lt8bNTm9pzFbNBScsH45x-FfeXA0ZFe", "g8")
	// crawlVideo("PLCd8j6ZYo0lYaVvkI0VXe9rwIgsYw78dG", "g8")
	// crawlVideo("PLzS2cTFGpDyzQhiRWI4l-zwWzkBRApUhx", "g8")
	// crawlVideo("PLb86fQcyLH4S8Mn0IZ2jjTWbC9MIpCQBn", "g8")
	// crawlVideo("PL-0HTRU3ATXC_7cHfgF6oSpvcJOBOcr0w", "g8")
	// crawlVideo("PLCd8j6ZYo0lZTVJ_6GWxKiSU-eywgoHjY", "g8")
	// crawlVideo("PL-0HTRU3ATXB1soXOaQ9l6L65dzfuPnK4", "g8")
	// crawlVideo("PL5kxOdJ5P409eg4ueMOi83OKLLXABb-aD", "g8")
	// crawlVideo("PL7lt8bNTm9pyxKadt4tR1Mt4ATp8Pk9_G", "g8")
	// crawlVideo("PLzS2cTFGpDywyMUXiA1xtW8hVk4WLIfb1", "g8")
	// crawlVideo("PLzS2cTFGpDywpTh_5uG0kkGOHR-q8IaX4", "g8")
	// crawlVideo("PL5kxOdJ5P40-h6SDfZCaoB6qDdrW1Nw6A", "g8")
	// crawlVideo("PL5kxOdJ5P40_ESh5uticYLEHCdEDY5ta8", "g8")
	// crawlVideo("PL5kxOdJ5P40_ESh5uticYLEHCdEDY5ta8", "g8")
	// crawlVideo("PLqcCOZrA1NA305ValbMOhrgijtmsQd-W8", "g8")
	// crawlVideo("PLYmriOlxq-mhdeQJ5Zcr5tw9nMkvdts1O", "g8")
	// crawlVideo("PLDHr_ecbSve65I5YduK8fGRjw2fA6QaZ7", "g8")
	// crawlVideo("PLBox31Ou7rrEzSiUyVvgI2wz5XtIxo1gp", "g8")
	// crawlVideo("PLxUdsmN71otBUYLix55ttAuzFWOO-dEiY", "g8")
	// crawlVideo("PLVZVOefqB0pAp20nHyFGT9P4X28MUA8uY", "g8")
	// crawlVideo("PLxUdsmN71otBSprJWbJJAeoMcuu4FJNmg", "g8")
	// crawlVideo("PLCd8j6ZYo0laWogGY2fbYk3HPwfr2l42b", "g8")
	// crawlVideo("PLHlw36ksrsNkVx2HDPWS1gFN3WNc4fFuE", "g8")
	// crawlVideo("PLxUdsmN71otBI690IKodc0U6zIZJCkTyW", "g8")
	// crawlVideo("PLCd8j6ZYo0lbaRFxN-5eMDXKAxbG5Bgqh", "g8")
	// crawlVideo("PLMzRq608THSMnLV5SGnMiwCYXG3tXAnnP", "g8")
	///lớp 9
	// crawlVideo("PLAFUuLnPbl6qffB1HPgBiumCCaQKdDFKY", "g9")
	// crawlVideo("PLSfUYrkcPcrvImpEF9-zEmldhlkZBkwuk", "g9")
	// crawlVideo("PLo3yP85_LULkmjQmkVtF-0gVGaRIwfjQi", "g9")
	// crawlVideo("PLL538aEph8_P8KpBlSyxgdezo9GT0IULp", "g9")
	// crawlVideo("PLA33KQsrdG3i55G3ob10MNGS0qIp_6tve", "g9")
	// crawlVideo("PLD2UOexoPAIH1c_fm9Je0_iDLUYr1Ju7H", "g9")
	// crawlVideo("PLzS2cTFGpDyy8nKJZfraRKladtraoYqOZ", "g9")
	// crawlVideo("PLphsmGG5Xfk8onpMmmMmBQO9U_qc0LhRN", "g9")
	// crawlVideo("PLjP9o7pgsmOUrT_amJMK9c-Val3GKhp3Y", "g9")
	// crawlVideo("PLxARqqwLS9WPD47lxX70dquOAk7TYyUUe", "g9")
	// crawlVideo("PLuMGE4ip9nTUGY2tNtWdq9GaEcw_1hjlb", "g9")
	// crawlVideo("PLAFUuLnPbl6qw_-2sX63DIhkrkEjXRR-y", "g9")
	// crawlVideo("PLER9IqxRQFWk6ExAlT_Caghs901LK8B4M", "g9")
	// crawlVideo("PL8x0QuuuxE-qXKEj1bdt9Bywo8ENeu06s", "g9")
	// crawlVideo("PLoXStX_pVftvjyz8qpq1crm1RuyYtkncH", "g9")
	// crawlVideo("PLD2UOexoPAIEaAAiQdN-L4oTDC-gsWNHj", "g9")
	// crawlVideo("PLAFUuLnPbl6rfWVqahdQAGSy0IAPyyt44", "g9")
	// crawlVideo("PLrkH4NRsM18OmmLk4Zo8foQVAFrqWjpjn", "g9")
	// crawlVideo("PLAFUuLnPbl6r-N8V31FqyFJP-Q9KgXx4A", "g9")
	// crawlVideo("PLxUdsmN71otCyfx6EEuD-eelFFtkZREN1", "g9")
	// crawlVideo("PLuMGE4ip9nTVQctlGYNKHN5P-5IRqZ77p", "g9")
	// crawlVideo("PLzS2cTFGpDyxRiy85wsnArqhvMRRHF_5f", "g9")
	// crawlVideo("PLYJOmh9gUZxlwce-YorN2da31KYnZEw6Y", "g9")
	// crawlVideo("PLCd8j6ZYo0la5gCPVGEo25oT6guxb9j5o", "g9")
	// crawlVideo("PLhKgNgqfK6qkjy7uN1n3mueb6QpRsCAhV", "g9")
	// crawlVideo("PL5kxOdJ5P40_g3Vyvjes4E0CBeMo6gqy5", "g9")
	// crawlVideo("PL19OTEQOJ53AepuKxCrRhUFGyn-AxQCTQ", "g9")
	// crawlVideo("PLD2UOexoPAIGCt6WE6D9Jp_1JWbQFFooJ", "g9")
	// crawlVideo("PLCd8j6ZYo0lY0i-wtos1VSPMF4FSuiZ4l", "g9")
	// crawlVideo("PLAFUuLnPbl6odWNlM0F9WhFPzKTZ2sfKo", "g9")
	// crawlVideo("PLYmriOlxq-mipR_ng-yzc00e2s2T3PYfW", "g9")
	// crawlVideo("PLzS2cTFGpDyxjW0zzoSNEKzlnLhGOuQ_s", "g9")
	// crawlVideo("PLMzRq608THSOlznpqDmZ2Y-DEDLuWCHSS", "g9")
	// crawlVideo("PLhKgNgqfK6qmGvs0igo92dlWslnuHRSVk", "g9")
	// crawlVideo("PLAFUuLnPbl6ql-NYXQahNpOkF9h4_NmWV", "g9")
	// crawlVideo("PLCd8j6ZYo0lbCKtj_jlWIA0EtH4sl4Cpj", "g9")
	// crawlVideo("PLzS2cTFGpDyyRXqu1jmple4AL4MfOs-q9", "g9")
	// crawlVideo("PLYJOmh9gUZxlok4YSSO6GprJt6puTTs-B", "g9")
	// crawlVideo("PLCd8j6ZYo0lYF5eYFmGpWyCN6exKchTwM", "g9")
	// crawlVideo("PL-qYbpxhuvCNIAqOax4Yr9ABN3cXx4fGf", "g9")
	// crawlVideo("PLCd8j6ZYo0lYbMTfkglZ6vj6VNsLP2vej", "g9")
	// crawlVideo("PLDHr_ecbSve5g7Sr4o8BLYZwJ91aOahhd", "g9")
	// crawlVideo("PLb86fQcyLH4RSXgzOpmWkVrdxv_Y8FqZT", "g9")
	// crawlVideo("PL5kxOdJ5P40-_zAWu0cDFpp-Aq-FSF38P", "g9")
	// crawlVideo("PLCd8j6ZYo0lYOu1QxTGRO2Xl1UnBk3EjW", "g9")
	// crawlVideo("PLoXStX_pVftv3BGVHvVpoXTQwHznDa4Ki", "g9")
	// crawlVideo("PL-0HTRU3ATXCDfeQFRk7JGdj1XNNdmVd-", "g9")
	// crawlVideo("PLCDT0UhUBPKO7JOw10tyxDOcDOiqc1N06", "g9")
	// crawlVideo("PLpJ_lT2rOwX1Dm9qvbKHA_9w7Q56dTgdG", "g9")
	// crawlVideo("PLXmeri-X8nVwUBWiz6CYjW_lFe4t2WA42", "g9")
	// crawlVideo("PLCd8j6ZYo0lY8ZFrhrAyzCzuo5x9YIrAm", "g9")
	// crawlVideo("PLDHr_ecbSve7neyouEH8Ptxod-AElpVaO", "g9")
	// crawlVideo("PL5kxOdJ5P40_1qWZdturGAe1PT5LVb1JY", "g9")
	// crawlVideo("PLzS2cTFGpDyyEtR-sIZTO1sIMbuJh6dmG", "g9")
	// crawlVideo("PL5kxOdJ5P409owUFVtfteFkxyJzFgAajR", "g9")
	// crawlVideo("PLhKgNgqfK6qlz8tbSSLiDVxhXCWY-q9g1", "g9")
	// crawlVideo("PLVZVOefqB0pBMPIy0a9L7sCztiXTh78P6", "g9")
	// crawlVideo("PLxUdsmN71otD-8c7wuPQ7wH4g38TNs2Vx", "g9")
	// crawlVideo("PLHlw36ksrsNlvBE-_yvlOWfXnNoCFaCgJ", "g9")
	// crawlVideo("PLxUdsmN71otBBRA8tasl2nMzHO-mO-Xbt", "g9")
	// crawlVideo("PLCd8j6ZYo0lbsBu_HLaCRncx-MGQvXwTJ", "g9")
	// crawlVideo("PLCd8j6ZYo0lYcWcuhR5N8v-Hw6oePUaRT", "g9")
	// crawlVideo("PL5kxOdJ5P40_Z0YgVgXxdjMweXq2t4Syr", "g9")
	// crawlVideo("PLhKgNgqfK6qkyJKAPQs5jStMGnLXjk30z", "g9")
	// crawlVideo("PLMzRq608THSNsKTy3gZ_MT41WJNMy5qrc", "g9")
	///ôn thi vào 10
	// crawlVideo("PLoXStX_pVftv7n5kNIiHR19JtXnOjRQ9V", "hsp")
	// crawlVideo("PL6BUF2OKNu1icjN2GewGc3VeHUHzH-yw7", "hsp")
	// crawlVideo("PLH5GyRVRwLCkVCbVxj-npQA3oNWoIj9bL", "hsp")
	// crawlVideo("PLZDTpmIiMAB8eCGZcLzF2gnfTbrw-cUU6", "hsp")
	// crawlVideo("PLR8LRnAndu4XH6AqQvye4sLtCTqBx-Hzx", "hsp")
	// crawlVideo("PL59YYxN58hcpQssdRhGm1xwQavHVqa7VD", "hsp")
	// crawlVideo("PLvTopE-zkuS7zDKdR-rmi1xP8v455Rf8l", "hsp")
	// crawlVideo("PLYgxl915x_r4PN0gwapzJK7JBvtO0NJf3", "hsp")
	// crawlVideo("PLY2M5jMMnXZdSmQtCnEPN4djbIfCjjnBO", "hsp")
	// crawlVideo("PLbLoggqJA9RGcSCQLhMWJAqcgi_FMnoj7", "hsp")
	// crawlVideo("PLDmMT65XgJD3SN1wIATEIK77kynWSFO-n", "hsp")
	// crawlVideo("PLidK1VaE4fKcvIvEvrAjFdFwvlLkTo1Vm", "hsp")
	// crawlVideo("PLZDTpmIiMAB8K8MvI9-4Q4OhB0O5U_GZK", "hsp")
	// crawlVideo("PLrwTMD00nrbDJP3P3W1GDXViZ2PkFeHvm", "hsp")
	// crawlVideo("PLNfDeaDOX8En83Co4BGMJ9A0JjcBscmut", "hsp")
	// crawlVideo("PLrdAluHlWG-DsIIHG-nfM0v2BVf8WyOkj", "hsp")
	// crawlVideo("PL2RzJLUvXc9JIp6Voe8SbQKxFaDksy9yu", "hsp")
	// crawlVideo("PLY2M5jMMnXZek8ZWEX5jyGd3d1R9Ttd75", "hsp")
	// crawlVideo("PL0et9Zydy571sfCbhN5BFP_tHLntohtw8", "hsp")
	// crawlVideo("PLIdtNiuDnsHCjl6LhQOrbdukkVt6qLXXr", "hsp")
	// crawlVideo("PLL538aEph8_P8KpBlSyxgdezo9GT0IULp", "hsp")
	// crawlVideo("PLD2UOexoPAIERwsJYL370yroEJAo8I1Nh", "hsp")
	// crawlVideo("PLD2UOexoPAIH1c_fm9Je0_iDLUYr1Ju7H", "hsp")
	// crawlVideo("PLXPgM84qwrZBKbqYfgvt2Aqqj0nfmURGa", "hsp")
	// crawlVideo("PLD2UOexoPAIEaAAiQdN-L4oTDC-gsWNHj", "hsp")
	// crawlVideo("PLD2UOexoPAIEaAAiQdN-L4oTDC-gsWNHj", "hsp")
	// crawlVideo("PLDHr_ecbSve5Jq5OCNeiPFz2gHrZbZXYo", "hsp")
	// crawlVideo("PLNuAYEJTtDMG5Yy72cP19kEWfRgExbLSW", "hsp")
	// crawlVideo("PLh5DVm-2geoo6wAXNCs6K3vy4D1Yq4BDT", "hsp")
	// crawlVideo("PLFUQUxTlUn61xIDqwIPbnjT35D-epF8yz", "hsp")
	// crawlVideo("PLCd8j6ZYo0la5gCPVGEo25oT6guxb9j5o", "hsp")
	// crawlVideo("PLfvc8A_Qkz_dyjLJUhY_Z_OmVPRf1YaQJ", "hsp")
	// crawlVideo("PLoXStX_pVftvjyz8qpq1crm1RuyYtkncH", "hsp")
	// crawlVideo("PLPnkRD761zEi-4aiAcxyJOw87c6Nkx3X3", "hsp")
	// crawlVideo("PL8x0QuuuxE-pyZCMtCldiiyGSEVB8YSzR", "hsp")
	// crawlVideo("PLb86fQcyLH4QsjBIa5BhMfkACJ1lPdJQ3", "hsp")
	// crawlVideo("PLoXStX_pVftv3BGVHvVpoXTQwHznDa4Ki", "hsp")
	// crawlVideo("PLoXStX_pVftv7n5kNIiHR19JtXnOjRQ9V", "hsp")
	///lớp 10
	// crawlVideo("PLBJb6uyJz5CGakAU2rwfoQAN4qnvpYnbi", "g10")
	// crawlVideo("PLuOh1vF0MkoG-ky-6DjO8_T-VKrx0Ad6u", "g10")
	// crawlVideo("PL7GwbLv3hi0i4N-fKFpexgfUZH4sVzKi8", "g10")
	// crawlVideo("PLNAnUfQYPEiGki4onXS02TBiuE9pnejSq", "g10")
	// crawlVideo("PL5iP0O_mGKihZF762jaMh27vOzTNW1Frl", "g10")
	// crawlVideo("PLLLv3emqnRa443V8sf5CUtfP_4-pZZ1qZ", "g10")
	// crawlVideo("PLba6cs1ag-qC1JFOjEE6SnsWySnY2tQyx", "g10")
	// crawlVideo("PLwJ4-6t2XvxAHHmgqOy-6o3-ypU1OHd7m", "g10")
	// crawlVideo("PLtxzK191_ZG1PbF1SUdD7VMmsuIxBGUPH", "g10")
	// crawlVideo("PLSz3Y-f1lM6uhxvBzCHfuQAp8O7PNYFt2", "g10")
	// crawlVideo("PLYJOmh9gUZxm3VckiurLXG1YimMNBJrLD", "g10")
	// crawlVideo("PL3ReErKGeKp88gYoObON7EQQq4dvO1yYs", "g10")
	// crawlVideo("PL5kxOdJ5P40-F126IhFXwN02G5ZqT9Z_J", "g10")
	// crawlVideo("PLXmeri-X8nVx-4MKKv5jRREmH_VZst-5Z", "g10")
	// crawlVideo("PLXmeri-X8nVydyYNKUcwsWZ3tETLtvPD5", "g10")
	// crawlVideo("PLCd8j6ZYo0lY3Bo42wbV_ABRpp0sbPsxN", "g10")
	// crawlVideo("PLEeAYV9Zmy-i7B55JewbdtKkTpMYykBcc", "g10")
	// crawlVideo("PLoXStX_pVftt_Djcvgq1UAmtHOKWoNC62", "g10")
	// crawlVideo("PLE8UM7QG5LzZUwub1I_gqW609KS16vK-b", "g10")
	// crawlVideo("PLXYp7Odn5ED_7kBWf-odmqFIpMLUbRZI_", "g10")
	// crawlVideo("PLoXStX_pVftt_Djcvgq1UAmtHOKWoNC62", "g10")
	// crawlVideo("PLXmeri-X8nVyZHG2_M3gC7_7lSZQQ2f9L", "g10")
	// crawlVideo("PLCd8j6ZYo0lZ9ASyPwNuWJ3VqmLavAiR7", "g10")
	// crawlVideo("PLXmeri-X8nVy6Kai_ameATXhWt-YAzHPY", "g10")
	// crawlVideo("PLXmeri-X8nVxgCl0mx1MtE6fwOk1ePFen", "g10")
	// crawlVideo("PLMzRq608THSOgROvlIyW8myVglaQ5iPfr", "g10")
	// crawlVideo("PLCd8j6ZYo0laruhgRmxC8zrfd5PY6Ne4u", "g10")
	// crawlVideo("PLX8bW_S_7PfnH8NjKw2K9YJiYPVhXC_H0", "g10")
	// crawlVideo("PLCd8j6ZYo0lYG_l7ZNJQEcEyhwtna1Kgf", "g10")
	// crawlVideo("PLXmeri-X8nVxGI0dYuO5bsc7UnBM4ZzLP", "g10")
	// crawlVideo("PLNAnUfQYPEiFnzsrdkmI2UDrEtAhBqlR_", "g10")
	///lớp 11
	// crawlVideo("PL0RcIuyecdhKjCsaNSfIkhyhIOmYGbiyu", "g11")
	// crawlVideo("PLuOh1vF0MkoGok-YCEjos4m7ZcJlwWL3u", "g11")
	// crawlVideo("PLK7zW26pwiHdM4Lz3OVtXVMZ-ZPJhhqgN", "g11")
	// crawlVideo("PLOz0SKVB63i1qWfpntwQ66SuSaJCUfs0T", "g11")
	// crawlVideo("PLOz0SKVB63i3A9N_0N160VMTFn7opYv6Q", "g11")
	// crawlVideo("PLuOh1vF0MkoF55yH_Y4kYQnGjkLx_5Q8b", "g11")
	// crawlVideo("PLZ0pWPk70Jv6poMIICFsrAmA3L6yLnL71", "g11")
	// crawlVideo("PL3ReErKGeKp-LHFVMpz9SF__Blg5EIYTF", "g11")
	// crawlVideo("PL0k2ozWJRpevDuJ0Q0kfG82Z6W8gBX8YW", "g11")
	// crawlVideo("PLaETZ7zfHfmKzDNoLJydhxC3bQ7JDifrW", "g11")
	// crawlVideo("PLPbA4EIPRHVPwjwpXtiWOlPrgMB93_hgR", "g11")
	// crawlVideo("PLuOh1vF0MkoFQsV6xSgpxDQevUJPTNKuH", "g11")
	// crawlVideo("PL-j4Nnf9AaHcdh7HjHiAw_NbQUOLCW8PG", "g11")
	// crawlVideo("PL7GwbLv3hi0jcJpKBSyzXeYOp2IvB3iGY", "g11")
	// crawlVideo("PL90p-tTX2LFqi2dvDghPd_YWYR6vLkwbv", "g11")
	// crawlVideo("PLLLv3emqnRa45cY7z0jf58LnP70bcYRJK", "g11")
	// crawlVideo("PLXmeri-X8nVwP5YHs4C1FmP3vJVIGxUHo", "g11")
	// crawlVideo("PLR3yj1XHVfSOKDMY5C0UUKrG1bhNQV5zO", "g11")
	// crawlVideo("PLCd8j6ZYo0lZQORsvi6INVSqDdNNNQRST", "g11")
	// crawlVideo("PLXmeri-X8nVx4B0H5Lpo1MvTAc_k0vpB1", "g11")
	// crawlVideo("PLCd8j6ZYo0laYOPASTepBqmTv1YCleMpG", "g11")
	// crawlVideo("PLba6cs1ag-qCrBQJo-sOtBkSWj3tkybO-", "g11")
	// crawlVideo("PLxNSRfTNtmiVn23tHeGmVcbiIXjVC1jRJ", "g11")
	// crawlVideo("PLxWx8us5g3p0mTdxW4CeKpk2Gr_jrtE9j", "g11")
	// crawlVideo("PLXmeri-X8nVx7nST89b2OMLXa0ff2hjCt", "g11")
	// crawlVideo("PLMzRq608THSOSlyng6b_BIzeZheR5DKlR", "g11")
	// crawlVideo("PLtxzK191_ZG1xXsLeJINMWYvK2LEdNEh6", "g11")
	// crawlVideo("PLeap9CEYpXjCyHy7SCqut50_Ne-WucSxY", "g11")
	// crawlVideo("PL3VcW1ai5zFAddxKgs8v2m8B5Cr_g89iU", "g11")
	// crawlVideo("PLXmeri-X8nVyIJSGgNkjnnETxPG-YgBLH", "g11")
	// crawlVideo("PLOz0SKVB63i3yz3OK2BvVBuUUMDrWVi-z", "g11")
	// crawlVideo("PLXmeri-X8nVyL4NaNgLu3XTBESFTBLhnu", "g11")
	// crawlVideo("PLCd8j6ZYo0lbaqU2G1uj7ncJJPq5huN3K", "g11")
	// crawlVideo("PLXmeri-X8nVxTeleONxDQ1pjR6H_oQbmw", "g11")
	// crawlVideo("PLX8bW_S_7Pfl_JBawTMkWDS32KiHGttN6", "g11")
	///lớp 12
	// crawlVideo("PLopuqf3UNTUdWxxYQC53GE_6ySsWHLoIM", "g12")
	// crawlVideo("PLOz0SKVB63i1KqjWUp26aoDbucFIY3p2O", "g12")
	// crawlVideo("PLyqgnSdPFZnAxw_2MjHvlhHTqAD0gZ1wB", "g12")
	// crawlVideo("PLuOh1vF0MkoGGWvPFctYJMHfrVXKaCgBP", "g12")
	// crawlVideo("PLXmeri-X8nVz9J0mTjEIViHhaNeK4mq7r", "g12")
	// crawlVideo("PL7xNhYHXNSNpULI3wU4517DxUA1Vq3NiX", "g12")
	// crawlVideo("PLq0mRSDfY0BCTfqY4ZTuGWGBf5S9XiGF3", "g12")
	// crawlVideo("PLIE2mz4PxLRKtl8Ppsoi6ZumetndzU3ac", "g12")
	// crawlVideo("PLCd8j6ZYo0lYo1jKiUbF2DvuqEputUHAH", "g12")
	// crawlVideo("PLQEjO_JbzKWUFm4uXlG_mNc1Ca3IfDOyN", "g12")
	// crawlVideo("PLbWzgQJ1hbxCKCnF1XaGLKtSHBFI55oV8", "g12")
	// crawlVideo("PLSz3Y-f1lM6vO-mdOJlo0vcY0tWi2Du7d", "g12")
	// crawlVideo("PLe9Mjl61ROzARUCWwMwPddSolmKDnSsUB", "g12")
	// crawlVideo("PLuOh1vF0MkoGLwKMvx_I9zwE7xBKJvLI6", "g12")
	// crawlVideo("PLR5Ie1493WKl2soePNm3WePIb1dfLScBd", "g12")
	// crawlVideo("PLba6cs1ag-qB0Y-N6nAxtHgsyvl7ywlx6", "g12")
	// crawlVideo("PLnZjL6hpKYAwJlgloRFN6cqCbzbY4_9Ku", "g12")
	// crawlVideo("PLba6cs1ag-qCIWc8fEXts3PFkS_HxDvTx", "g12")
	// crawlVideo("PLCd8j6ZYo0lZr9VMeeb-lWFrbTp41h5mA", "g12")
	// crawlVideo("PLXYp7Odn5ED87zwdrTer49vbphlsTGU_9", "g12")
	// crawlVideo("PLq0mRSDfY0BC6FabIJw9-MxhAfHPakZPs", "g12")
	// crawlVideo("PLo3yP85_LULnKwP_YTIZdr8tQJ0jhGXj1", "g12")
	// crawlVideo("PLOz0SKVB63i0mkhuw6tBhovIvmbAMETtg", "g12")
	// crawlVideo("PLXmeri-X8nVydIoUTujAYaV1uaeavkpjA", "g12")
	// crawlVideo("PLCd8j6ZYo0lYj4RKiKpZOUAbyYAcEjOWS", "g12")
	// crawlVideo("PLXmeri-X8nVylZXWzIh9hsH1zIjS686OT", "g12")
	// crawlVideo("PLg7LaSBus8No3Ki5OTu6sHKpaUedSfYc1", "g12")
	// crawlVideo("PLXmeri-X8nVz44g73HUIKaN_atMmUc1Uj", "g12")
	// crawlVideo("PLXmeri-X8nVyQhoAJ9mxE05wOBQBCGUjg", "g12")
	// crawlVideo("PLXmeri-X8nVw2tGQ0yqasy-o_jogn2vJG", "g12")
	// crawlVideo("PLXmeri-X8nVzUbd9dFHY6E4o80kxAJz7v", "g12")
	// crawlVideo("PLXmeri-X8nVzh1d_AjoKoTqZI1D3iqU5q", "g12")
	// crawlVideo("PLCd8j6ZYo0lYJlV2hC3rsZjGTNyIKAHbs", "g12")
	// crawlVideo("PLe9Mjl61ROzBMleUTUpRdT0zsDUm-md1N", "g12")
	// crawlVideo("PLoXStX_pVfttsdWvf9n7qptxuRXybsKAZ", "g12")
	// crawlVideo("PLW0JdsZ4tzvDsx388NXOk-sMYokhRRNuH", "g12")
	// crawlVideo("PLMzRq608THSN1w-nqAd8URCekx0k4qEnX", "g12")
	// crawlVideo("PLXmeri-X8nVx0PHXb53cfAIS9eH60GkdK", "g12")
	// crawlVideo("PLVPfdl8gNB5np2eTfyAVH_Nq4Upw6NNKp", "g12")
	// crawlVideo("PLwJ4-6t2XvxAsQHWn_3plThG6WYi-LV3m", "g12")
	// crawlVideo("PLXmeri-X8nVwKZM86S_TmWS2UsvB_mi45", "g12")
	// crawlVideo("PLXYp7Odn5ED_IWlliBd985TMNmaWmjW4b", "g12")
	// crawlVideo("PLX8bW_S_7Pfnl4baLc5EG1R1Br1Jo7ICW", "g12")
	///ThptQG
	// crawlVideo("PLopuqf3UNTUdWxxYQC53GE_6ySsWHLoIM", "up")
	// crawlVideo("PLi0bzuaaHLJQhg3Nc-faZ5Plw-eXlG04h", "up")
	// crawlVideo("PLcY1SLXqNi8TwyAbAyxnu-BJcMD4n7Y0W", "up")
	// crawlVideo("PLHdjApfoSgTtstARtol8woECb3G5jJMuO", "up")
	// crawlVideo("PLbWzgQJ1hbxCxmEjjrJV0GUrmRIaJr5CV", "up")
	// crawlVideo("PLoaoacxbeQ-XpcvUAqQHFnF4HTaTTfpSA", "up")
	// crawlVideo("PLtxzK191_ZG2auu-C571s2jfGuJHX7yNp", "up")
	// crawlVideo("PLHdjApfoSgTs0H6kGmAdxacROspjEnaur", "up")
	// crawlVideo("PLq0mRSDfY0BCN6DwMs0DAQepP_3Hk7QMO", "up")
	// crawlVideo("PLo3yP85_LULnKwP_YTIZdr8tQJ0jhGXj1", "up")
	// crawlVideo("PLq0mRSDfY0BAYuPXMeSZ65cVcL9qiCh-t", "up")
	// crawlVideo("PLq0mRSDfY0BDs1619vkAgsTWT7gWCYFl_", "up")
	// crawlVideo("PLW0JdsZ4tzvDsx388NXOk-sMYokhRRNuH", "up")
	// crawlVideo("PLV40fEjmhD80pRQq308sEhitzPyNJncg8", "up")
	// crawlVideo("PLeG4Zq1DsiQtTjs2F2dHOCZ53Oi6RZkMo", "up")
	// crawlVideo("PLXYp7Odn5ED8jj_ROzHTVt5H4NNYZtY66", "up")
}

func crawlVideo(playlist string, gr string) {
	lastModified, err := vRep.GetLastModifiedPlaylist(playlist)
	if err != nil {
		fmt.Println("Lỗi lấy last modified playlist, tạm dừng:", err)
		return
	}
	fmt.Println("Last modified playlist:", lastModified)

	fmt.Println("Đang lấy thông tin playlist:", playlist)

	// --flat-playlist + --print để lấy nhanh video + playlist name
	cmd := exec.Command("yt-dlp", "--flat-playlist", "--print", "%(playlist_title)s|%(playlist_id)s|%(id)s|%(title)s", playlist)
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(output), "\n")
	var videos []models.Video
	var newLastModified time.Time
	var playlistTitle string
	var playlistID string

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "|")
		if len(parts) < 4 {
			continue
		}

		playlistTitle = normalizeUTF8(parts[0])
		playlistID = parts[1]
		videoID := parts[2]
		title := normalizeUTF8(parts[3])

		if title == "- YouTube" || strings.TrimSpace(title) == "" {
			fmt.Println("Bỏ qua video private:", videoID)
			continue
		}

		url := "https://www.youtube.com/watch?v=" + videoID
		uploadDate, err := getYouTubeUploadDate(url)
		if err != nil {
			fmt.Println("Lỗi lấy ngày đăng:", err)
			continue
		}

		if uploadDate.Before(lastModified) {
			continue
		}

		if uploadDate.After(newLastModified) {
			newLastModified = uploadDate
		}

		video := models.Video{
			Id:           videoID,
			Title:        title,
			URL:          url,
			Grade:        gr,
			LastModified: uploadDate,
			Playlist:     playlist,
		}
		fmt.Printf("\n----- Thêm video mới: %s\nTên: %s\nLớp: %s\nUrl: %s\n Cập nhật: %s", video.Id, video.Title, video.Grade,video.URL,video.LastModified)
		videos = append(videos, video)
	}

	fmt.Printf("Tổng số video mới cần thêm: %d\n", len(videos))
	if len(videos) == 0 {
		return
	}

	pl := models.Playlist{
		Id:           playlistID,
		Title:        playlistTitle,
		Thumbnail:    videos[0].Id,
		Grade:        gr,
		Count:        len(videos),
		LastModified: newLastModified,
	}

	fmt.Printf("\n>>>>>Playlist mới: %s\nTitle: %s\nCount: %s\nLastmodified: %s", pl.Id, pl.Title, pl.Count, pl.LastModified)

	//Upload videos và playlist nếu cần
	vRep.UploadVideos(videos)
	vRep.UploadPlaylist(pl)
}

func normalizeUTF8(s string) string {
	// Nếu s không hợp lệ UTF-8 thì chuyển đổi từng rune
	if utf8.ValidString(s) {
		return s
	}
	// Loại bỏ ký tự không hợp lệ
	return strings.ToValidUTF8(s, "")
}

func getYouTubeUploadDate(videoURL string) (time.Time, error) {
    cmd := exec.Command("yt-dlp", "-j", videoURL)
    output, err := cmd.Output()
    if err != nil {
        return time.Time{}, err
    }

    var data struct {
        UploadDate      string `json:"upload_date"`      // "YYYYMMDD"
        UploadTimestamp int64  `json:"upload_timestamp"` // epoch seconds
    }

    if err := json.Unmarshal(output, &data); err != nil {
        return time.Time{}, err
    }

    if data.UploadTimestamp != 0 {
        return time.Unix(data.UploadTimestamp, 0), nil
    }

    if data.UploadDate != "" {
        t, err := time.Parse("20060102", data.UploadDate)
        if err != nil {
            return time.Time{}, err
        }
        return t, nil
    }

    return time.Time{}, fmt.Errorf("không tìm thấy upload_timestamp hoặc upload_date")
}



func getYouTubeID(videoURL string) string {
	re := regexp.MustCompile(`(?:v=|youtu\.be/)([A-Za-z0-9_-]{11})`)
	matches := re.FindStringSubmatch(videoURL)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func getYouTubeTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return "", fmt.Errorf("title not found")
		case html.StartTagToken:
			t := tokenizer.Token()
			if t.Data == "title" {
				tokenizer.Next()
				title := strings.TrimSpace(tokenizer.Token().Data)
				return strings.TrimSuffix(title, " - YouTube"), nil
			}
		}
	}
}
