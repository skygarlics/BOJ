def bench(*fns, rep=10000, show_output=False):
    """Benchmarking function with stdio redirection"""
    import io
    import sys
    import timeit
    
    if len(fns) == 0:
        raise ValueError("At least one function must be provided")
    for fn in fns:
        if not callable(fn):
            fn_name = fn.__name__ if hasattr(fn, '__name__') else str(fn)
            raise ValueError(f"{fn_name} is not a callable")
        
    input_data = sys.stdin.read()

    times = {}
    if show_output: print("=== Outputs ===")
    for fn in fns:
        fn_name = fn.__name__ if hasattr(fn, '__name__') else str(fn)
        def wrapper():
            sys.stdin = io.StringIO(input_data)
            sys.stdout = io.StringIO()
            fn()
        total_time = timeit.timeit(wrapper, number=rep)
        times[fn_name] = total_time

        if show_output:
            output = sys.stdout.getvalue()
            sys.stdout = sys.__stdout__
            print(f"{fn_name} output:\n{output}")

    sys.stdout = sys.__stdout__
    print("=== Benchmarks ===")
    for k, v in times.items():
        print(f"{k}: {v:.6f}")