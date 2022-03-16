<script>
  import { push } from 'svelte-spa-router'

  export let color = 'blue'
  export let link
  export let target
</script>

{#if link}
  {#if link.startsWith('http')}
    <a class="btn {color}" href={link} {target}><slot/></a>
  {:else}
    <a class="btn {color}" on:click={() => push(link)} href={link} {target}><slot/></a> 
  {/if}
{:else}
  <a class="btn {color}" on:click><slot/></a>
{/if}


<style>
  .btn {
		height: 32px;

		display: flex;
		justify-content: center;
		align-items: center;
		gap: 8px;

		padding: 0px 8px;
		font-size: 12px;
		font-weight: 600;

		border-radius: 8px;
		border: none;

    color: var(--complementary-white);

		cursor: pointer;

		transition: 0.3s;
	}

	.btn:hover {
		opacity: 0.9;
	}

  .blue {
		background-color: var(--primary-blue);
  }

  .green {
		background-color: var(--primary-green);
  }
</style>