/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		'./templates/**/*.{html,js,templ,go}', // Make sure to include Go templates
		'./templates/common/**/*.{html,templ,go}', // Templates inside common
		'./templates/components/**/*.{html,templ,go}' // Templates inside components
	],
	theme: {
		extend: {}
	},
	plugins: []
}
