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
    background-color: var(--complementary-white-5);
    transition: 0.2s ease-in-out;
  }

  .normal:hover {
    box-shadow: 0px 0px 2px 0px rgba(255, 255, 255, 0.623);
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