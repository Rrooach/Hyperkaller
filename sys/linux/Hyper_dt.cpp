include <linux/fcntl.h>
include <linux/ioctl.h>

resource fddt[int32]: -1
open$IOCTL_PRIVCMD_HYPERCALLdt(file ptr[in, string["/proc/xen/privcmd"]]) fddt

ioctl$IOCTL_PRIVCMD_HYPERCALLdt0(fd fddt, ioctl_privcmd_hypecall0 const[0x305000], hypercall ptr[in, my_hypercalldt0])
ioctl$IOCTL_PRIVCMD_HYPERCALLdt1(fd fddt, ioctl_privcmd_hypecall0 const[0x305000], hypercall ptr[in, my_hypercalldt1])
 
close$IOCTL_PRIVCMD_HYPERCALLdt(fd fdmem)

my_hypercalldt0 {
	__HYPERVISOR_hypercall_type	const[0, int64]
	args				array[int64, 5]
}

my_hypercalldt1 {
	__HYPERVISOR_hypercall_type	const[10, int64]
	args				array[int64, 5]
}
 