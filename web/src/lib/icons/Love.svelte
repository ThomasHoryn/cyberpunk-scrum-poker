<script lang="ts">
    	let {size} = $props();
</script>

<style>
.heart-pulse {
  animation: pulse 1.5s infinite;
}

.heart-icon-container {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

@keyframes pulse {
  0% { transform: scale(1); }
  50% { transform: scale(1.2); }
  100% { transform: scale(1); }
}

</style>

<svg xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" 
class="heart-pulse"
width={size}
height={size} x="0" y="0" 
viewBox="0 0 512 512" 
style="enable-background:new 0 0 512 512" xml:space="preserve"
><g>
    <path d="M402.837 288.631h-32.631V256h32.631zM239.685 158.108h-32.631v-32.631h32.631zM304.946 158.108h-32.631v-32.631h32.631zM272.315 190.739h-32.631v-32.631h32.631z" fill="#ff0000" opacity="1" data-original="#000000" class=""></path><path d="M337.576 125.478h-32.631V92.847h32.631zM370.207 125.478h-32.631V92.847h32.631zM174.424 125.478h-32.631V92.847h32.631z" fill="#ff0000" opacity="1" data-original="#000000" class=""></path><path d="M207.054 125.478h-32.631V92.847h32.631zM435.468 190.739h-32.631v-32.631h32.631zM402.837 158.108h-32.631v-32.631h32.631z" fill="#ff0000" opacity="1" data-original="#000000" class=""></path><path d="M435.468 223.369h-32.631v-32.631h32.631zM435.468 256h-32.631v-32.631h32.631zM109.162 190.739H76.531v-32.631h32.631zM141.794 158.108h-32.631v-32.631h32.631z" fill="#ff0000" opacity="1" data-original="#000000" class=""></path><path d="M109.162 223.369H76.531v-32.631h32.631zM109.162 256H76.531v-32.631h32.631zM304.946 386.522h-32.631v-32.631h32.631z" fill="#ff0000" opacity="1" data-original="#000000" class=""></path><path d="M337.576 353.892h-32.631v-32.631h32.631zM370.207 321.261h-32.631V288.63h32.631zM207.054 353.892h-32.631v-32.631h32.631z" fill="#ff0000" opacity="1" data-original="#000000" class=""></path><path d="M174.424 321.261h-32.631V288.63h32.631zM272.315 419.153h-32.631v-32.631h32.631zM239.685 386.522h-32.631v-32.631h32.631z" fill="#ff0000" opacity="1" data-original="#000000" class=""></path><path d="M141.794 288.631h-32.631V256h32.631z" fill="#ff0000" opacity="1" data-original="#000000"></path></g></svg>