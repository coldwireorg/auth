<script>
  export let type = 'p'
  export let color
</script>

{#if type === 'h1'}
  <h1 style="color: {color};"><slot/></h1>
{:else if type === 'h2'}
  <h2 style="color: {color};"><slot/></h2>
{:else if type === 'h3'}
  <h3 style="color: {color};"><slot/></h3>
{:else if type === 'p'}
  <p style="color: {color};"><slot/></p>
{/if}

<style>
  h1, h2, h3, h4, h5, p {
    margin: 0;
  }

  h1 {
    font-weight: 600;
    font-size: 36px;
  }

  h2 {
    font-weight: 600;
    font-size: 18px;
  }

  h3 {
    font-weight: 600;
    font-size: 12px;
  }

  p {
    font-weight: 600;
    font-size: 10px;
  }
</style>