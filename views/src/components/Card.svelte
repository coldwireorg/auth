<script>
  export let type = 'normal'
  export let style
  export let hover = false
</script>

<div {style} class="card {type} {hover ? 'hover': ''}">
  <slot/>
</div>

<style>
  .card {
    display: flex;
    justify-content: flex-start;

    padding: 16px;

    border-radius: 8px;

    height: min-content;
  }

  .normal {
    background-color: var(--complementary-white-5);
    transition: 0.2s ease-in-out;
  }

  .green {
    border: none;
    background-color: var(--primary-green);
  }

  .green-25 {
    border: none;
    background-color: var(--primary-green-25);
  }

  .clear {
    background: var(--complementary-gray-3);
  }

  .hover:hover {
    box-shadow: 0px 0px 2.7px 0px rgba(255, 255, 255, 0.623);
  }
</style>