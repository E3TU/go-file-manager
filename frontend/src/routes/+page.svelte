<script>
    import Icon from "@iconify/svelte";
</script>

<div class="container">
    <h1>Go file manager</h1>
    <div class="file-upload">
        <p>Upload a file</p>
        <Icon class="upload-icon" icon="solar:upload-outline" width="100" height="100" />
    </div>
</div>

<style>
    .container{
        display: flex;
        align-items: center;
        flex-direction: column;
        height: 100vh;
        width: 100%;
    }
    h1{
        color: var(--text-primary);
        margin-top: 5rem;
        font-size: 3rem;
    }
    .file-upload{
        display: flex;
        flex-direction: column;
        align-items: center;
        height: 20rem;
        width: 20rem;
        background-color: var(--gray);
        box-shadow: rgba(0, 0, 0, 0.16) 0px 10px 36px 0px, rgba(0, 0, 0, 0.06) 0px 0px 0px 1px;
        border-radius: 16px;
        margin-top: 2rem;
    }
    p{
        color: var(--text-primary);
        padding: 1rem;
    }
    :global(.upload-icon) {
        color: var(--text-primary);
        margin-top: 3rem;
    }
</style>