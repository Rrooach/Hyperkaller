#include <err.h>
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <xenctrl.h>
#include "xencov.h"
static xc_interface *xch = NULL;

int cov_sysctl(int op, struct xen_sysctl *sysctl, struct xc_hypercall_buffer *buf, uint32_t buf_size)
{
    DECLARE_HYPERCALL_BUFFER_ARGUMENT(buf);

    memset(sysctl, 0, sizeof(*sysctl));
    sysctl->cmd = XEN_SYSCTL_coverage_op;

    sysctl->u.coverage_op.cmd = op;
    sysctl->u.coverage_op.size = buf_size;
    set_xen_guest_handle(sysctl->u.coverage_op.buffer, buf);

    return xc_sysctl(xch, sysctl);
}

int cov_read(void)
{
    xch = xc_interface_open(NULL, NULL, 0); 
    struct xen_sysctl sys;
    uint32_t total_len;
    DECLARE_HYPERCALL_BUFFER(uint8_t, p);
    const char *fn = "/dev/cov";
    FILE *f; 
    if (cov_sysctl(XEN_SYSCTL_COVERAGE_get_size, &sys, NULL, 0) < 0)
    {
        err(1, "getting total length");
        return 1;
    } 
    total_len = sys.u.coverage_op.size; 
    /* Shouldn't exceed a few hundred kilobytes */
    if (total_len > 8u * 1024u * 1024u)
    {
        errx(1, "gcov data too big %u bytes\n", total_len);
        return 1;
    } 
    p = xc_hypercall_buffer_alloc(xch, p, total_len);
    if (!p)
    {
        err(1, "allocating buffer");
        return 1;
    } 
    memset(p, 0, total_len);
    if (cov_sysctl(XEN_SYSCTL_COVERAGE_read, &sys, HYPERCALL_BUFFER(p),
                    total_len) < 0)
    {
        err(1, "getting gcov data");
        return 1;   
    } 
    if (!strcmp(fn, "-"))
        f = stdout;
    else
        f = fopen(fn, "w"); 
    if (!f)
    {
        err(1, "opening output file");
        return 1;
    } 

    if (fwrite(p, 1, total_len, f) != total_len)
    {
        err(1, "writing gcov data to file");
        return 1;
    } 
    if (f != stdout)
        fclose(f); 
    xc_hypercall_buffer_free(xch, p);
    xc_interface_close(xch);
    return 0;
}

// static void cov_reset(void)
// {
//     xch = xc_interface_open(NULL, NULL, 0);    
//     struct xen_sysctl sys;

//     if (cov_sysctl(XEN_SYSCTL_COVERAGE_reset, &sys, NULL, 0) < 0)
//         err(1, "resetting gcov information");
//     xc_interface_close(xch);
// }
 