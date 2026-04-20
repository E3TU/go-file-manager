<script lang="ts">
	import Icon from '@iconify/svelte';
	import Navbar from '../components/Navbar.svelte';
	import { uploadFile, listFiles, deleteFile, type UploadedFile } from '../lib/api';
	import { onMount } from 'svelte';

	let files = $state<UploadedFile[]>([]);
	let isDragging = $state(false);
	let inputEl: HTMLInputElement | null = null;
	let isUploading = $state(false);
	let error = $state<string | null>(null);

	async function loadFiles() {
		try {
			const res = await listFiles();
			files = res.files;
		} catch (e) {
			console.error('Failed to load files:', e);
		}
	}

	async function onDrop(e: DragEvent) {
		e.preventDefault();
		isDragging = false;
		if (e.dataTransfer?.files?.length) {
			await uploadFiles(Array.from(e.dataTransfer.files));
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

	async function onInputChange(e: Event) {
		const target = e.target as HTMLInputElement;
		if (target.files?.length) {
			await uploadFiles(Array.from(target.files));
			target.value = '';
		}
	}

	async function uploadFiles(fileList: File[]) {
		isUploading = true;
		error = null;
		try {
			for (const file of fileList) {
				const uploaded = await uploadFile(file);
				files = [...files, uploaded];
			}
		} catch (e) {
			error = e instanceof Error ? e.message : 'Upload failed';
		} finally {
			isUploading = false;
		}
	}

	async function removeFile(file: UploadedFile) {
		try {
			await deleteFile(file.id);
			files = files.filter((f) => f.id !== file.id);
		} catch (e) {
			error = e instanceof Error ? e.message : 'Delete failed';
		}
	}

	function formatSize(bytes: number): string {
		if (bytes < 1024) return bytes + ' B';
		if (bytes < 1024 * 1024) return (bytes / 1024).toFixed(1) + ' KB';
		return (bytes / (1024 * 1024)).toFixed(1) + ' MB';
	}

	onMount(() => {
		loadFiles();
	});
</script>

<Navbar />
<div class="container">
	{#if error}
		<p class="error">{error}</p>
	{/if}
	<div class="file-upload">
		<div
			class:dragging={isDragging}
			ondrop={onDrop}
			ondragover={onDragOver}
			ondragleave={onDragLeave}
		>
			<Icon class="upload-icon" icon="solar:upload-outline" width="64" height="64" />
			<p>Drag and drop files here to upload<br /> or</p>
			<button onclick={openFileManager} class="choose-file-btn" disabled={isUploading}>
				{isUploading ? 'Uploading...' : 'Choose file'}
			</button>
			<input onchange={onInputChange} bind:this={inputEl} class="file-input" type="file" multiple />
		</div>
	</div>
	<ul class="files">
		{#each files as file}
			<li>
				<span class="file-name">{file.name}</span>
				<span class="file-size">{formatSize(file.sizeOriginal)}</span>
				<a href={file.downloadUrl} target="_blank" class="download-link">
					<Icon icon="solar:download-outline" width="20" height="20" />
				</a>
				<button onclick={() => removeFile(file)} class="delete-btn">
					<Icon icon="solar:trash-bin-2-outline" width="20" height="20" />
				</button>
			</li>
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
	.error {
		color: #ff4444;
		margin-top: 1rem;
	}
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
	.choose-file-btn:hover:not(:disabled) {
		background-color: var(--blue);
		transition: 0.5s;
		cursor: pointer;
	}
	.choose-file-btn:disabled {
		opacity: 0.5;
		cursor: not-allowed;
	}
	.file-input {
		display: none;
	}
	.files {
		color: var(--text-primary);
		list-style: none;
		padding-top: 2rem;
		font-size: 1rem;
		width: 25rem;
	}
	.files li {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		padding: 1rem;
		background: var(--gray);
		border-radius: 8px;
		margin-bottom: 0.5rem;
		box-shadow:
			rgba(0, 0, 0, 0.16) 0px 3px 6px,
			rgba(0, 0, 0, 0.23) 0px 3px 6px;
	}
	.file-name {
		flex: 1;
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}
	.file-size {
		color: #888;
		font-size: 0.9rem;
	}
	.download-link,
	.delete-btn {
		color: var(--text-primary);
		background: none;
		border: none;
		cursor: pointer;
		padding: 0.25rem;
		display: flex;
		align-items: center;
		transition: 0.5s;
	}
	.delete-btn:hover {
		color: #ff4444;
		transition: 0.5s;
	}
	.download-link:hover {
		color: var(--blue);
		transition: 0.5s;
	}
</style>
