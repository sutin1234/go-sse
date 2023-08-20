<script>
	import { onMount } from 'svelte';
	let time = '';

	onMount(() => {
		const evtSource = new EventSource('http://localhost:3500/event');
		evtSource.onmessage = (evt) => {
			console.log('onmessage ', evt)
			time = evt.data;
		};

		evtSource.onerror = (evt) => {
			console.log('onErr ', evt)
		}
	});

	async function getTime(){
		const res = await fetch("http://localhost:3500/time")
		if(res.status !== 200){
			console.log('Cloud not connect to server')
		}
	}
</script>

<section>
	<h1>Svelte kit + SEE Go</h1>
	time : { time }
	<div>
		<button on:click={getTime}>get Time</button>
	</div>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}

	h1 {
		width: 100%;
	}

	.welcome {
		display: block;
		position: relative;
		width: 100%;
		height: 0;
		padding: 0 0 calc(100% * 495 / 2048) 0;
	}

	.welcome img {
		position: absolute;
		width: 100%;
		height: 100%;
		top: 0;
		display: block;
	}
</style>
