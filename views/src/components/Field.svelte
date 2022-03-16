<script>
  export let type;
	export let placeholder;
	export let value;
</script>

{#if type == 'password'}
  <input on:submit class="input" type="password" bind:value {placeholder} />
{:else if type == 'text'}
  <input on:submit class="input" type="text" bind:value {placeholder} />
{/if}

<style>
  input {
    height: 32px;
		width: auto;

		background-color: var(--complementary-gray-1);
		border: 1px solid var(--complementary-white-25);
		border-radius: 8px;

		font-size: 12px;

		padding: 0 15px 0 15px;

		color: var(--color-white);
  }
</style>