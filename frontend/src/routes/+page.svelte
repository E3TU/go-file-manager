<script lang="ts">
	import Icon from '@iconify/svelte';
	import Navbar from '../components/Navbar.svelte';

	// let fileInput: any;

	let files = $state<File[]>([]);

	let isDragging = $state(false);
	let inputEl: HTMLInputElement | null = null;

	function onDrop(e: DragEvent) {
		e.preventDefault();
		isDragging = false;
		if (e.dataTransfer?.files?.length) {
			files = [...files, ...Array.from(e.dataTransfer.files)];
		}
	}

	function onDragOver(e: DragEvent) {
		e.preventDefault();
		isDragging = true;
	}

	function onDragLeave() {
		isDragging = false;
	}

	function openFileManager() {
		inputEl?.click();
	}

    function onInputChange(e: Event) {
        const target = e.target as HTMLInputElement;
        if (target.files?.length) {
            files = [...files, ...Array.from(target.files)];
            target.value = "";
        }
    }
</script>

<Navbar />
<div class="container">
	<!-- <h1>Go file manager</h1> -->
	<div class="file-upload">
		<div
			class:dragging={isDragging}
			ondrop={onDrop}
			ondragover={onDragOver}
			ondragleave={onDragLeave}
		>
			<Icon class="upload-icon" icon="solar:upload-outline" width="64" height="64" />
			<p>Drag and drop files here to upload<br /> or</p>
			<button onclick={openFileManager} class="choose-file-btn">Choose file</button>
			<input onchange={onInputChange} bind:this={inputEl} class="file-input" type="file" multiple />
		</div>
	</div>
	<ul class="files">
		{#each files as file}
			<li>{file.name} ({file.size} bytes)</li>
		{/each}
	</ul>
</div>

<style>
	.container {
		display: flex;
		align-items: center;
		flex-direction: column;
		height: 100vh;
		width: 100%;
	}
	/* h1 {
		color: var(--text-primary);
		margin-top: 5rem;
		font-size: 3rem;
	} */
	.file-upload {
		display: flex;
		flex-direction: column;
		align-items: center;
		height: 15rem;
		width: 25rem;
		background-color: var(--gray);
		box-shadow:
			rgba(0, 0, 0, 0.16) 0px 10px 36px 0px,
			rgba(0, 0, 0, 0.06) 0px 0px 0px 1px;
		border-radius: 16px;
		margin-top: 5rem;
		border: 2px dashed var(--text-primary);
		text-align: center;
	}
	p {
		color: var(--text-primary);
		padding: 1rem;
		font-size: 1.1rem;
	}
	:global(.upload-icon) {
		margin-top: 2rem;
		color: var(--blue);
		transition: 0.3s;
	}
	.choose-file-btn {
		font-size: 1rem;
		font-family: inherit;
		border: 2px solid var(--blue);
		border-radius: 8px;
		padding: 0.5rem 1rem;
		background-color: transparent;
		color: var(--text-primary);
		transition: 0.5s;
	}
	.choose-file-btn:hover {
		background-color: var(--blue);
		transition: 0.5s;
		cursor: pointer;
	}
	.file-input {
		display: none;
	}
	.files{
		color: var(--text-primary);
		list-style: none;
		padding-top: 2rem;
		font-size: 1.5rem;
	}
</style>
