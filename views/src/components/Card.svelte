<script>
  export let type = 'normal'
  export let style
</script>

<div {style} class="card {type}">
  <slot/>
</div>

<style>
  .card {
    display: flex;
    justify-content: flex-start;

    padding: 16px;

    border-radius: 8px;
  }

  .normal {
    border: 2px solid var(--complementary-white-5);
    background-color: var(--complementary-white-5);
  }

  .green {
    border: none;
    background-color: var(--primary-green);
  }

  .green-25 {
    border: none;
    background-color: var(--primary-green-25);
  }
</style>