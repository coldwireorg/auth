<script>
  import { push } from 'svelte-spa-router'

  export let link
</script>

{#if link.startsWith('http')}
  <a href={link} target="_blank"><slot/></a>
{:else}
  <span on:click={() => push(link)} class="link"><slot /></span>  
{/if}

<style>
  .link {
    cursor: pointer;
  }
</style>